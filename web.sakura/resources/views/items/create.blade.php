@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Create an Item</h3>
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
                <label for="image" class="col-form-label">Main Image</label>
                {{ Form::file('image') }}
                <p class="form-text">
                    Try and find a somewhat decent quality image for the main image!<br>
                    <strong>Do not upload the default.png image. Leave this blank if you don't have a main image.</strong>
                </p>

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
            <div class="form-label-group">
                {{ Form::text('english_name', null, ['placeholder' => 'English Name', 'class' => 'form-control', 'required' => true]) }}
                <label for="english_name">English Name <span class="text-danger">*</span></label>

                <p class="form-text">
                    An english or romanized version of this item's name. <span title="This will be used to identify the item in the search index. If a name is a duplicate, a -0 or -1 (etc). will be added to the generated URL for this item." data-placement="top" data-toggle="tooltip" class="far fa-info-circle"></span>
                </p>
            </div>

            <div class="form-label-group">
                {{ Form::text('foreign_name', null, ['placeholder' => 'Foreign Name', 'class' => 'form-control', 'required' => true]) }}
                <label for="foreign_name">Foreign Name <span class="text-danger">*</span></label>

                <p class="form-text">The non-english version of this item's name. Usually the original.</p>
            </div>

            <div class="row">
                <div class="col-lg-4">
                    <div class="form-label-group">
                        <select name="year" id="year" class="form-control form-control-chosen" required>
                            @foreach (array_reverse(range(1990, date('Y'))) as $year)
                                <option value="{{ $year }}" @if ($year == old('year', date('Y'))) selected @endif>{{ $year }}</option>
                            @endforeach
                        </select>
                        <label for="year">Release Year <span class="text-danger">*</span></label>

                        <p class="form-text">The release year.</p>
                    </div>
                </div>

                <div class="col-lg-8">
                    <div class="form-label-group">
                        {{ Form::text('product_number', null, ['placeholder' => 'Product Number', 'class' => 'form-control']) }}
                        <label for="product_number">Product Number</label>

                        <p class="form-text">The seller's product number for this item.</p>
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col-lg-4">
                    <div class="form-label-group">
                        <select id="currency" name="currency" class="form-control form-control-chosen" required>
                            @foreach ($currencies as $code => $currency)
                                <option value="{{ $code }}" @if ($code === old('currency')) selected @endif>{{ $currency }}</option>
                            @endforeach
                        </select>
                        <label for="currency">Currency <span class="text-danger">*</span></label>

                        <p class="form-text">Use the release currency.</p>
                    </div>
                </div>
                <div class="col-lg-8">
                    <div class="form-label-group">
                        {{ Form::text('price', null, ['placeholder' => '1000', 'class' => 'form-control','required' => true]) }}
                        <label for="price">Item Price <span class="text-danger">*</span></label>

                        <p class="form-text">Prices should be in whole numbers! If unknown, use 0.</p>
                    </div>
                </div>
            </div>

            <div class="form-label-group">
                {{ Form::textarea('notes', null, ['placeholder' => 'Item Notes. Note that this no longer supports HTML/bbcode, but may in the future!', 'class' => 'form-control']) }}
                <label for="notes">Item Notes</label>
            </div>

            <h2>Attributes</h2>
            <div class="form-group">
                <p class="form-text">
                    Select a button below to add that particular attribute to this item.
                </p>
                <div class="col-lg-12" style="margin-bottom: 20px">
                    @foreach ($attributes as $attribute)
                        <button
                            type="button" class="btn btn-sm btn-outline-primary"
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
                        <div class="form-label-group" id="attribute-{{ $attribute->id }}" style="display: none">
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
                            <label for="attribute.{{ $attribute->slug }}">{{ $attribute->name }}</label>
                        </div>
                    @endforeach
                </div>
            </div>


    <div class="row">
        <div class="col-lg-12 text-center">
            <button type="submit" class="btn btn-outline-primary btn-lg">Save as Draft</button>
        </div>
    </div>

    {{ Form::close() }}
</div>
@endsection

@section('script')
    <script>
        $(function () {
            $('[data-type="attribute.button"]').click(function (event) {
                event.preventDefault();

                var $button = $(this),
                    data = $button.data(),
                    $attribute = $('#attribute-' + data.id);

                console.log(data);
                console.log($attribute);

                if (data.clicked) {
                    $button.addClass('btn-outline-primary').removeClass('btn-primary');
                    $button.data('clicked', 0);

                    $attribute.hide();
                } else {
                    $button.removeClass('btn-outline-primary').addClass('btn-primary');
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

            $('[data-toggle="tooltip"]').tooltip();
        });
    </script>

</div>
@endsection
