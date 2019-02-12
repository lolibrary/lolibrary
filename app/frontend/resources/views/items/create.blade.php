@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Create an Item<</h3>
        </div>
    </div>
    <h1>Create an Item</h1>
    @if ($errors->any())
    <div class="alert alert-danger">
        <ul>
            @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
        </ul>
    </div>
@endif

            {{ Form::open(['route' => 'items.store', 'files' => true, 'method' => 'put']) }}
            <div class="row">
                <div class="col-lg-4">
                    <div class="form-group">
                        <label for="image" class="col-form-label">Main Image <span class="text-danger">*</span></label>
                        {{ Form::file('image') }}
                        <p class="form-text">Try and find a somewhat decent quality image for the main image!</p>
                    </div>

                    <hr>

                    <div class="form-group">
                        <label for="image" class="col-form-label">Additional Images</label>
                        {{ Form::file('images[]', ['id' => 'additional-image']) }}

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
                                <option value="{{ $feature->id }}" @if (in_array($feature->id, old('features', []), true)) selected @endif>{{ $feature->name }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="colors" class="col-form-label">Colorways</label>

                        <select id="colors" name="colors[]" class="form-control form-control-chosen" multiple>
                            @foreach ($colors as $color)
                                <option value="{{ $color->id }}" @if (in_array($color->id, old('colors', []), true)) selected @endif>{{ $color->name }}</option>
                            @endforeach
                        </select>
                    </div>
                    
                    <div class="form-group">
                        <label for="tags" class="col-form-label">Tags</label>

                        <select name="tags[]" id="tags" class="form-control form-control-chosen" multiple>
                            @foreach ($tags as $tag)
                                <option value="{{ $tag->id }}" @if (in_array($tag->id, old('tags', []), true)) selected @endif>{{ $tag->slug }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="brand" class="col-form-label">Brand <span class="text-danger">*</span></label>

                        <select id="brand" name="brand" class="form-control form-control-chosen" required>
                            @foreach ($brands as $brand)
                                <option value="{{ $brand->id }}" @if ($brand->id === old('brand')) selected @endif>{{ $brand->name }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="category" class="col-form-label">Category <span class="text-danger">*</span></label>

                        <select id="category" name="category" class="form-control form-control-chosen" required>
                            @foreach ($categories as $category)
                                <option value="{{ $category->id }}" @if ($category->id === old('category')) selected @endif>{{ $category->name }}</option>
                            @endforeach
                        </select>
                    </div>
                </div>
                <div class="col-lg-8">
                    <div class="form-group">
                        <label for="english_name" class="col-form-label">English Name <span class="text-danger">*</span></label>
                        {{ Form::text('english_name', null, ['placeholder' => 'English Name', 'class' => 'form-control', 'required' => true]) }}
                        <p class="form-text">
                            An english or romanized version of this item's name. This will be used to identify the item in the search index.
                            <br>
                            If a name is a duplicate, a -0 or -1 (etc). will be added to the generated URL for this item.
                        </p>
                    </div>

                    <div class="form-group">
                        <label for="foreign_name" class="col-form-label">Foreign Name <span class="text-danger">*</span></label>
                        {{ Form::text('foreign_name', null, ['placeholder' => 'Foreign Name', 'class' => 'form-control', 'required' => true]) }}

                        <p class="form-text">The non-english version of this item's name. Usually the original.</p>
                    </div>

                    <div class="form-group">
                        <label for="product_number" class="col-form-label">Product Number</label>
                        {{ Form::text('product_number', null, ['placeholder' => 'Product Number', 'class' => 'form-control']) }}

                        <p class="form-txt">The seller's product number for this item.</p>
                    </div>

                    <div class="form-group">
                        <label for="year" class="col-form-label">Release Year <span class="text-danger">*</span></label>

                        <select name="year" id="year" class="form-control form-control-chosen" required>
                            @foreach (array_reverse(range(1990, date('Y') + 3)) as $year)
                                <option value="{{ $year }}" @if ($year == old('year', date('Y'))) selected @endif>{{ $year }}</option>
                            @endforeach
                        </select>
                    </div>

                    <div class="row">
                        <div class="col-lg-4">
                            <div class="form-group">
                                <label for="currency" class="col-form-label">Currency <span class="text-danger">*</span></label>

                                <select id="currency" name="currency" class="form-control form-control-chosen" required>
                                    @foreach ($currencies as $code => $currency)
                                        <option value="{{ $code }}" @if ($code === old('currency')) selected @endif>{{ $currency }}</option>
                                    @endforeach
                                </select>
                                <p class="form-text">Use the release currency.</p>
                            </div>
                        </div>
                        <div class="col-lg-8">
                            <div class="form-group">
                                <label for="price" class="col-form-label">Item Price <span class="text-danger">*</span></label>

                                {{ Form::text('price', null, ['placeholder' => '1000', 'class' => 'form-control','required' => true]) }}
                                <p class="form-text">Prices should be in whole numbers! (hint: you can use 0 here but please try not to!)</p>
                            </div>
                        </div>
                    </div>

                    <div class="form-group">
                        <label for="notes" class="col-form-label">Item Notes</label>
                        {{ Form::textarea('notes', null, ['placeholder' => 'Item Notes. Note that this no longer supports HTML/bbcode, but may in the future!', 'class' => 'form-control']) }}
                    </div>

                    <div class="form-group">
                        <label class="ccol-form-label">Attributes</label>
                        <p class="form-text">
                            Select a button below to add that particular attribute to this item.
                        </p>
                        <div class="col-lg-12" style="margin-bottom: 20px">
                            @foreach ($attributes as $attribute)
                                <button
                                    type="button" class="btn btn-sm btn-secondary"
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
                                        null,
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


            <div class="row">
                <div class="col-lg-12 text-center">
                    <button type="submit" class="btn btn-primary btn-lg">Save as Draft</button>
                </div>
            </div>

            {{ Form::close() }}
        </div>
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

</div>
@endsection
