@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Edit an Item</h1>
        </div>
    </div>

    @if ($errors->any())
    <div class="alert alert-danger">
        <ul>
            @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
        </ul>
    </div>
    @endif

    @if ($item->published())
        <div class="row">
            <div class="col-lg-4">
                    <ul class="list-group">
                        <li class="list-group-item">Status: <span class="badge badge-success">PUBLISHED</span></li>
                        <li class="list-group-item">
                            Publisher: {{ $item->publisher->name or "Unknown"}}
                            @if ($item->publisher && $item->publisher->is(auth()->user()))
                                (You)
                            @endif
                        </li>
                        <li class="list-group-item">Published: {{ $item->published_at->format('d M Y \a\t H:i:s') }}</li>
                    </ul>
                </div>

            @if ($item->submitter && $item->submitter->is(auth()->user()) && auth()->user()->lolibrarian())
                <div class="col-lg-8">
                    <div class="card text-white bg-success">
                        <div class="card-body">
                            You are able to edit this (published) item.
                        </div>
                    </div>
                </div>
            @elseif (auth()->user()->senior())
                <div class="col-lg-8">
                    <div class="card text-white bg-success">
                        <div class="card-body">
                            You are able to edit this (published) item.
                        </div>
                    </div>
                </div>
            @else
                <div class="col-lg-8">
                    <div class="card text-white bg-danger">
                        <div class="card-body">
                            You are unable to edit this item now that it has been published.
                        </div>
                    </div>
                </div>
            @endif
        </div>
    @endif

    <div class="col-sm p-2 px-4">
            {{ Form::model($item, ['route' => ['items.update', $item], 'files' => true, 'method' => 'put']) }}
            <div class="row">
                <div class="col-lg-4">
                    <div class="form-group">
                        <label for="image" class="col-form-label">Replace Main Image</label>

                        <div>
                            <a href="{{ $item->image->url }}">
                                <img src="{{ $item->image->thumbnail_url }}" alt="" style="height: 80px; max-width: 100%">
                            </a>
                        </div>

                        {{ Form::file('image') }}
                    </div>

                    <hr>

                    <div class="form-group">
                        <label for="image" class="col-form-label">Additional Images</label>

                        @foreach ($item->images as $image)
                            <div class="row text-center">
                                <div class="col-6">
                                    <a href="{{ $image->url }}">
                                        <img src="{{ $image->thumbnail_url }}" alt="" style="height: 80px; max-width: 100%">
                                    </a>
                                </div>
                                <div class="col-6">
                                    <a onclick="event.preventDefault(); $('#delete-image-{{ $image->id }}').submit();" class="btn btn-sm btn-danger">
                                        <i class="fal fa-fw fa-trash"></i> 
                                        delete
                                    </a>
                                </div>
                            </div>
                        @endforeach

                        {{ Form::file('images[]') }}

                        <div id="additional-container"></div>

                        <button
                            type="button"
                            class="btn btn-xs btn-primary"
                            onclick="$('<input name=\'images[]\' type=\'file\'>').appendTo($('#additional-container'))"
                        >add another image</button>
                    </div>

                    <hr>

                    <div class="form-group">
                        <label for="features" class="col-form-label">Features</label>

                        <select id="features" name="features[]" class="form-control form-control-chosen" multiple>
                            @foreach ($features as $feature)
                                <option value="{{ $feature->id }}" @if (in_array($feature->id, old('features', $item->features->pluck('id')->all()), true)) selected @endif>{{ $feature->name }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="colors" class="col-form-label">Colorways</label>

                        <select id="colors" name="colors[]" class="form-control form-control-chosen" multiple>
                            @foreach ($colors as $color)
                                <option value="{{ $color->id }}" @if (in_array($color->id, old('colors', $item->colors->pluck('id')->all()), true)) selected @endif>{{ $color->name }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="tags">Tags</label>

                        <select name="tags[]" id="tags" class="form-control form-control-chosen" multiple>
                            @foreach ($tags as $tag)
                                <option value="{{ $tag->id }}" @if (in_array($tag->id, old('tags', $item->tags->pluck('id')->all()), true)) selected @endif>{{ $tag->slug }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="brand" class="col-form-label">Brand <span class="text-danger">*</span></label>

                        <select id="brand" name="brand" class="form-control form-control-chosen" @if ($item->published()) disabled @else name="brand" @endif>
                            @foreach ($brands as $brand)
                                <option value="{{ $brand->id }}" @if ($brand->id === old('brand', $item->brand->id)) selected @endif>{{ $brand->name }}</option>
                            @endforeach
                        </select>
                        @if ($item->published())
                            You cannot change the brand of an item once it is published.
                        @endif
                    </div>

                    <div class="form-group">
                        <label for="category" class="col-form-label">Category <span class="text-danger">*</span></label>

                        <select id="category" class="form-control form-control-chosen" @if ($item->published()) disabled @else name="category" @endif>
                            @foreach ($categories as $category)
                                <option value="{{ $category->id }}" @if ($category->id === old('category', $item->category->id)) selected @endif>{{ $category->name }}</option>
                            @endforeach
                        </select>
                        @if ($item->published())
                            You cannot change the category of an item once it is published.
                        @endif
                    </div>
                </div>
                <div class="col-lg-8">
                    <div class="form-group">
                        <label for="english_name" class="col-form-label">English Name <span class="text-danger">*</span></label>
                        {{ Form::text('english_name', null, ['placeholder' => 'English Name', 'class' => 'form-control', 'required' => true]) }}
                        <p class="form-text">
                            An english or romanized version of this item's name. This will be used to identify the item in the search index.
                        </p>
                    </div>

                    <div class="form-group">
                        <label for="foreign_name" class="col-form-label">Foreign Name</label>
                        {{ Form::text('foreign_name', null, ['placeholder' => 'Foreign Name', 'class' => 'form-control', 'required' => true]) }}

                        <p class="form-text">The non-english version of this item's name. Usually the original.</p>
                    </div>

                    <div class="form-group">
                        <label for="product_number" class="col-form-label">Product Number</label>
                        {{ Form::text('product_number', null, ['placeholder' => 'Product Number', 'class' => 'form-control']) }}

                        <p class="form-text">The seller's product number for this item.</p>
                    </div>

                    <div class="form-group">
                        <label for="year">Release Year <span class="text-danger">*</span></label>

                        <select name="year" id="year" class="form-control form-control-chosen" required>
                            @foreach (array_reverse(range(1990, date('Y') + 3)) as $year)
                                <option value="{{ $year }}" @if ($year == old('year', $item->year)) selected @endif>{{ $year }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="row">
                        <div class="col-lg-4">
                            <div class="form-group">
                                <label for="currency" class="col-form-label">Currency <span class="text-danger">*</span></label>

                                <select id="currency" name="currency" class="form-control form-control-chosen" required>
                                    @foreach ($currencies as $code => $currency)
                                        <option value="{{ $code }}" @if ($code === old('currency', $item->currency)) selected @endif>{{ $currency }}</option>
                                    @endforeach
                                </select>
                                <p class="form-text">Use the release currency.</p>
                            </div>
                        </div>
                        <div class="col-lg-8">
                            <div class="form-group">
                                <label for="price" class="col-form-label">Item Price <span class="text-danger">*</span></label>

                                {{ Form::text('price', null, ['placeholder' => '1000', 'class' => 'form-control']) }}
                                <p class="form-text">Prices should be in whole numbers!</p>
                            </div>
                        </div>
                    </div>

                    <div class="form-group">
                        <label for="notes" class="col-form-label">Item Notes</label>
                        {{ Form::textarea('notes', null, ['placeholder' => 'Item Notes. Note that this no longer supports HTML/bbcode, but may in the future!', 'class' => 'form-control']) }}
                    </div>

                    <div class="form-group">
                        <label class="col-form-label">Attributes</label>
                        <p class="form-text">
                            Select a button below to add that particular attribute to this item.
                        </p>
                        <div class="col-lg-12" style="margin-bottom: 20px">
                            @foreach ($attributes as $attribute)
                                <button
                                    type="button" class="btn btn-sm btn-default"
                                    data-type="attribute.button"
                                    data-id="{{ $attribute->id }}"
                                    data-clicked="0"
                                    style="margin: 5px"
                                    id="attribute-button-{{ $attribute->id }}"
                                >
                                    {{ $attribute->name }}
                                </button>
                            @endforeach
                        </div>
                    </div>

                    <div class="row">
                        <div class="col-lg-12">
                            @foreach ($attributes as $attribute)
                                <div class="form-group" id="attribute-{{ $attribute->id }}" style="display: none">
                                    <label for="attribute.{{ $attribute->slug }}" class="col-form-label">{{ $attribute->name }}</label>
                                    {{ Form::text(
                                        "attributes[{$attribute->id}]",
                                        $item->attributes->contains($attribute)
                                            ? $item->attributes->find($attribute)->pivot->value
                                            : null,
                                        [
                                            'placeholder' => $attribute->name,
                                            'class' => 'form-control',
                                            'id' => "attribute.{$attribute->slug}",
                                            'data-type' => 'attribute.input',
                                            'data-id' => $attribute->id,
                                        ]
                                    ) }}
                                </div>
                            @endforeach
                        </div>
                    </div>
                </div>
            </div>


            <div class="row">
                <div class="col-lg-4 offset-lg-4 text-center">
                    @if ($item->published())
                        @if (
                            ($item->publisher && $item->publisher->is(auth()->user()) && auth()->user()->lolibrarian())
                            || auth()->user()->senior()
                        )
                            <button type="submit" class="btn btn-success btn-block btn-lg">
                                Update Item
                            </button>
                        @else
                            <button type="button" class="btn btn-success btn-block btn-lg" disabled>
                                Update Item
                            </button>
                        @endif
                    @else
                        <button type="submit" class="btn btn-primary btn-block btn-lg">
                            Save as Draft
                        </button>
                    @endif
                </div>
            </div>

            {{ Form::close() }}
        </div>

    @foreach ($item->images as $image)
        {{ Form::open([
            'route' => ['items.images.destroy', $item, $image],
            'method' => 'delete',
            'id' => 'delete-image-' . $image->id,
        ]) }}
        {{ Form::close() }}
    @endforeach

    {{ Form::open(['route' => ['items.destroy', $item], 'method' => 'delete', 'id' => 'delete-item']) }}
    {{ Form::close() }}
</div>
@endsection

@section('script')
    <script>
        $('[data-type="attribute.button"]').click(function (event) {
            event.preventDefault();
            var $button = $(this),
                data = $button.data(),
                $attribute = $('#attribute-' + data.id);
            console.log(data);
            console.log($attribute);
            if (data.clicked) {
                $button.addClass('btn-default').removeClass('btn-primary');
                $button.data('clicked', 0);
                $attribute.hide();
            } else {
                $button.removeClass('btn-default').addClass('btn-primary');
                $button.data('clicked', 1);
                $attribute.show();
            }
        });
        $('[data-type="attribute.input"]').each(function () {
            var $input = $(this),
                data = $input.data(),
                $button = $('#attribute-button-' + data.id);
            if ($input.val()) {
                $button.click();
            }
        });
    </script>

@endsection