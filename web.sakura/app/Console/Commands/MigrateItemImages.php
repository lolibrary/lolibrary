<?php

namespace App\Console\Commands;

use Exception;
use App\Models\Item;
use App\Models\Image;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\DB;

class MigrateItemImages extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'lolibrary:migrate-images {--status=all}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Migrates all main item images to flattened non-model files';

    /**
     * A list of all missing main images.
     * 
     * @var array
     */
    protected $missing = [];

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $query = $this->query($this->option('status'));

        $bar = $this->output->createProgressBar($query->count());

        foreach ($query->orderBy('id')->cursor() as $item) {
            if (! $this->mainImageActioned($item)) {
                $this->actionMainImage($item);
            }

            if (! $this->imagesActioned($item)) {
                $this->actionImages($item);
            }

            $bar->advance();
        }

        $bar->finish();
        $this->line('');

        $this->table(['Item ID', 'Image ID'], $this->missing);

        $this->line(count($this->missing) . ' items are missing images');
    }

    protected function query(string $status)
    {
        switch ($status) {
            case "all":
                return Item::query();
            case "published":
                return Item::where('status', Item::PUBLISHED);
            case "draft":
                return Item::where('status', Item::DRAFT);
        }

        if (is_numeric($status)) {
            return Item::where('status', $status);
        }

        throw new Exception("Cannot figure out what status was given: {$status}");
    }

    protected function mainImageActioned(Item $item): bool
    {
        return ! ($item->image === null || $item->image === '');
    }

    protected function imagesActioned(Item $item): bool
    {
        return !empty($item->images);
    }

    protected function actionMainImage(Item $item)
    {
        $image = Image::find($item->image_id);

        if ($image === null) {
            $this->missing[] = [$item->id, $item->image_id];

            return;
        }

        $item->image = "images/{$image->filename}";
        $item->save();
    }

    protected function actionImages(Item $item)
    {
        $pivot = collect(DB::select('select image_id from image_item where item_id = ?', [$item->id]))->map->image_id->all();

        $images = Image::whereIn('id', $pivot)->get();

        $json = [];

        foreach ($images as $image) {
            $json[] = [
                'key' => bin2hex(random_bytes(8)),
                'layout' => 'image',
                'attributes' => [
                    'image' => "images/{$image->filename}",
                ],
            ];
        }

        $item->images = $json;
        $item->save();
    }
}
