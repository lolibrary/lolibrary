<?php

namespace App\Notifications\ImageProcessing;

class InitialUploadFailed extends Notification implements ShouldQueue
{
    /**
     * The image that failed its initial upload.
     *
     * @param \App\Image
     */
    public $image;

    /**
     * Make a new notification instance.
     *
     * @param \App\Image $image
     */
    public function __construct(Image $image)
    {
        $this->image = $image;
    }

    public function via()
    {
        return ['slack'];
    }

    public function toSlack($notifiable)
    {
        $url = route('admin.images.show', $image);

        return (new SlackMessage)
            ->to($notifiable)
            ->title('Initial Image Upload Failed')
            ->markdown("[Image {$this->image->id}]($url) failed to upload.");
    }
}
