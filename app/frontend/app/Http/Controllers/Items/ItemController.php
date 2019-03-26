<?php

namespace App\Http\Controllers\Items;

use App\Models\Tag;
use App\Models\Item;
use App\Models\User;
use App\Models\Brand;
use App\Models\Color;
use App\Models\Image;
use App\Models\Feature;
use App\Models\Category;
use App\Models\Attribute;
use Illuminate\Http\Request;
use App\Http\Requests\Admin\{
    ItemStoreRequest, ItemUpdateRequest
};
use Illuminate\Http\UploadedFile;
use Illuminate\Support\Facades\DB;
use App\Http\Controllers\Controller;

class ItemController extends Controller
{
    /**
     * Show an item.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\Response
     */
    public function show(Item $item)
    {
        $item->load(Item::FULLY_LOAD);

        return view('items.show', compact('item'));
    }

    /**
     * Show a paginated list of items.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return redirect()->route('search');
    }

    /**
     * Update a user's wishlist.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\RedirectResponse
     */
    public function wishlist(Item $item)
    {
        $user = auth()->user();
        $attached = $user->updateWishlist($item);
        $status = $attached ? 'added' : 'removed';

        return back()->withStatus(trans("user.wishlist.{$status}", ['item' => str_limit($item->english_name, 28)]));
    }

    /**
     * Update a user's closet.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\RedirectResponse
     */
    public function closet(Item $item)
    {
        $user = auth()->user();
        $attached = $user->updateCloset($item);
        $status = $attached ? 'added' : 'removed';

        return back()->withStatus(trans("user.closet.{$status}", ['item' => str_limit($item->english_name, 28)]));
    }

    /**
     * Edit an item.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\RedirectResponse|\Illuminate\View\View
     */

    public function edit(Item $item)
    {
        $this->user = auth()->user();

        if ($item->published() && ! $this->user->senior()) {
            return back()->withErrors("Your level is not allowed to edit items once published!");
        }

        if ($item->submitter && ! $item->submitter->is($this->user) && ! $this->user->senior()) {
            return back()->withErrors("You are not allowed to edit someone else's submission.");
        }

        return view('items.edit', [
            'item' => $item->load(Item::FULLY_LOAD),
            'attributes' => Attribute::all(),
            'brands' => Brand::select(['id', 'name'])->get(),
            'categories' => Category::select(['id', 'name'])->get(),
            'features' => Feature::select(['id', 'name'])->get(),
            'colors' => Color::select(['id', 'name'])->get(),
            'tags' => Tag::select(['id', 'slug'])->get(),
            'currencies' => Item::CURRENCIES,
        ]);
    }

    /**
     * Create an item.
     *
     * @return \Illuminate\View\View
     */
    public function create()
    {
        return view('items.create', [
            'attributes' => Attribute::all(),
            'currencies' => Item::CURRENCIES,
            'brands' => Brand::select(['id', 'name'])->get(),
            'categories' => Category::select(['id', 'name'])->get(),
            'features' => Feature::select(['id', 'name'])->get(),
            'colors' => Color::select(['id', 'name'])->get(),
            'tags' => Tag::select(['id', 'slug'])->get(),
        ]);
    }

    /**
     * Create an item.
     *
     * @param \App\Http\Requests\Admin\ItemStoreRequest $request
     * @return \Illuminate\Http\Response
     */
    public function store(ItemStoreRequest $request)
    {
        DB::transaction(function () use ($request) {
            $brand = Brand::findOrFail($request->brand);
            $category = Category::findOrFail($request->category);

            // handle the main image.
            if ($request->image instanceof UploadedFile) {
                $image = Image::from($request->image);
            } else {
                $image = Image::default();
            }

            // handle the extra images
            $images = collect($request->images)->map(function (UploadedFile $file) {
                $image = Image::from($file);

                return $image->id;
            });

            $item = new Item($request->only([
                'english_name',
                'foreign_name',
                'notes',
                'year',
                'product_number',
                'price',
                'currency',
            ]));
            $item->brand()->associate($brand);
            $item->category()->associate($category);
            $item->image()->associate($image);
            $item->submitter()->associate(auth()->user());
            $item->status = Item::DRAFT;
            $item->slug = Item::slug($item);
            $item->save();
            // now we can add features, attributes and images.
            if ($images) {
                $item->images()->attach($images->all());
            }
            $item->features()->attach($request->features);
            $item->colors()->attach($request->colors);
            $item->tags()->attach($request->tags);
            $item->attributes()->attach(
                collect($request->input('attributes'))
                    ->filter()
                    ->map(function ($value, $key) {
                        return ["attribute_id" => $key, "value" => $value];
                    }
                    )
            );
        });

        return redirect()->route('items.show', $item);
    }

    /**
     * Create an item.
     *
     * @param \App\Http\Requests\Admin\ItemUpdateRequest $request
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\Response
     */
    public function update(ItemUpdateRequest $request, Item $item)
    {
        /** @var \App\Models\User $user */
        $user = auth()->user();
        if ($item->published()) {
            $allowed = ($user->is($item->submitter) && $user->lolibrarian()) || $user->senior();
        } else {
            $allowed = $user->is($item->submitter) || $user->senior();
        }
        if (! $allowed) {
            return back()->withErrors("Sorry, you can't do that!");
        }

        if ($item->draft()) {
            $brand = Brand::findOrFail($request->brand);
            $category = Category::findOrFail($request->category);
            $item->brand()->associate($brand);
            $item->category()->associate($category);
        }
        // handle the main image.
        if ($request->image instanceof UploadedFile) {
            $image = Image::from($request->image);
            $item->image()->associate($image);
        }
        // handle the extra images (can be done async)
        $images = collect($request->images)->map(function (UploadedFile $file) {
            $image = Image::from($file);
            return $image->id;
        });

        $item->fill($request->only([
            'english_name',
            'foreign_name',
            'notes',
            'year',
            'product_number',
            'price',
            'currency',
        ]));
        $item->save();
        // now we can add features, attributes and images.
        $item->images()->attach($images->all());
        $item->features()->sync($request->features);
        $item->colors()->sync($request->colors);
        $item->tags()->sync($request->tags);
        $item->attributes()->sync(
            collect($request->input('attributes'))
                ->filter()
                ->map(function ($value, $key) {
                    return ["attribute_id" => $key, "value" => $value];
                }
                )
        );

        return redirect()->route('items.show', $item);
    }

    /**
     * Publish an item and add it to the search index.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\Response
     */
    public function publish(Item $item)
    {
        if ($item->published()) {
            return back()->withErrors('That item is already published.');
        }
        /** @var \App\Models\User $user */
        $user = auth()->user();
        if (! $user->is($item->submitter) && ! $user->senior()) {
            // require senior to publish other's items.
            return back()->withErrors("You cannot publish another user's post with your access level!");
        }

        if (! $user->lolibrarian()) {
            return back()->withErrors("Sorry, you can't publish items with your role");
        }

        $item->publish();
        return back()->with('status', 'Item Published - It may take a few moments for it to appear in search results!');
    }

    /**
     * Delete an item.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\Response
     */
    public function destroy(Item $item)
    {
        $error = back()->withErrors("You don't have permission to do that.");
        /** @var \App\Models\User $user */
        $user = auth()->user();
        if ($item->published() && ! $user->admin()) {
            return $error;
        } elseif (! ($user->is($item->submitter) || $user->senior())) {
            return $error;
        }
        $item->delete();
        return redirect()
            ->route('items.index')
            ->with('status', 'Item deleted successfully');
    }

    /**
     * Delete an image on a post.
     *
     * @param \App\Models\Item $item
     * @param \App\Models\Image $image
     * @return \Illuminate\Http\RedirectResponse|\Illuminate\Http\Response
     */
    public function deleteImage(Item $item, Image $image)
    {
        // sanity check first
        if (! $item->images->contains($image)) {
            abort(404);
        }
        /** @var \App\Models\User $user */
        $user = auth()->user();

        if ($item->published()) {
            $allowed = ($user->is($item->submitter) && $user->lolibrarian()) || $user->senior();
        } else {
            $allowed = $user->is($item->submitter) || $user->senior();
        }

        if (! $allowed) {
            return back()->withErrors("You aren't allowed to do that!");
        }

        if ($image->id === uuid5('default')) {
            return back()->withErrors("You can't delete the default image.");
        }

        $image->delete();

        return back()->with('status', 'Image Deleted');
    }
}
