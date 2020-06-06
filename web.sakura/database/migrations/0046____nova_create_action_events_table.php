<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class NovaCreateActionEventsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::dropIfExists('action_events');
        Schema::create('action_events', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->char('batch_id', 36);
            $table->uuid('user_id')->index();
            $table->string('name');
            $table->string('actionable_type');
            $table->uuid('actionable_id');
            $table->string('target_type');
            $table->text('target_id');
            $table->string('model_type');
            $table->text('model_id')->nullable();
            $table->text('fields');
            $table->string('status', 25)->default('running');
            $table->text('exception');
            $table->timestamps();

            $table->index(['actionable_type', 'actionable_id']);
            $table->index(['batch_id', 'model_type', 'model_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('action_events');
    }
}
