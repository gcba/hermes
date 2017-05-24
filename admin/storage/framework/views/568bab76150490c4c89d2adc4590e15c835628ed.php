<?php $__env->startSection('css'); ?>
    <style>
        .panel-actions .voyager-trash {
            cursor: pointer;
        }

        .panel-actions .voyager-trash:hover {
            color: #e94542;
        }

        .panel hr {
            margin-bottom: 10px;
        }

        .panel {
            padding-bottom: 15px;
        }

        .sort-icons {
            font-size: 21px;
            color: #ccc;
            position: relative;
            cursor: pointer;
        }

        .sort-icons:hover {
            color: #37474F;
        }

        .voyager-sort-desc {
            margin-right: 10px;
        }

        .voyager-sort-asc {
            top: 10px;
        }

        .page-title {
            margin-bottom: 0;
        }

        .panel-title code {
            border-radius: 30px;
            padding: 5px 10px;
            font-size: 11px;
            border: 0;
            position: relative;
            top: -2px;
        }

        .new-setting {
            text-align: center;
            width: 100%;
            margin-top: 20px;
        }

        .new-setting .panel-title {
            margin: 0 auto;
            display: inline-block;
            color: #999fac;
            font-weight: lighter;
            font-size: 13px;
            background: #fff;
            width: auto;
            height: auto;
            position: relative;
            padding-right: 15px;
        }

        .new-setting hr {
            margin-bottom: 0;
            position: absolute;
            top: 7px;
            width: 96%;
            margin-left: 2%;
        }

        .new-setting .panel-title i {
            position: relative;
            top: 2px;
        }

        .new-settings-options {
            display: none;
            padding-bottom: 10px;
        }

        .new-settings-options label {
            margin-top: 13px;
        }

        .new-settings-options .alert {
            margin-bottom: 0;
        }

        #toggle_options {
            clear: both;
            float: right;
            font-size: 12px;
            position: relative;
            margin-top: 15px;
            margin-right: 5px;
            margin-bottom: 10px;
            cursor: pointer;
            z-index: 9;
            -webkit-touch-callout: none;
            -webkit-user-select: none;
            -khtml-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none;
            user-select: none;
        }

        .new-setting-btn {
            margin-right: 15px;
            position: relative;
            margin-bottom: 0;
            top: 5px;
        }

        .new-setting-btn i {
            position: relative;
            top: 2px;
        }

        textarea {
            min-height: 120px;
        }
        textarea.hidden{
            display:none;
        }
    </style>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('head'); ?>
    <script type="text/javascript" src="<?php echo e(voyager_asset('lib/js/jsonarea/jsonarea.min.js')); ?>"></script>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="voyager-settings"></i> Settings
    </h1>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>

    <div class="container-fluid">
        <?php echo $__env->make('voyager::alerts', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <?php if(config('voyager.show_dev_tips')): ?>
        <div class="alert alert-info">
            <strong>How To Use:</strong>
            <p>You can get the value of each setting anywhere on your site by calling <code>Voyager::setting('key')</code></p>
        </div>
        <?php endif; ?>
    </div>

    <div class="page-content container-fluid">
        <form action="<?php echo e(route('voyager.settings.update')); ?>" method="POST" enctype="multipart/form-data">
            <?php echo e(method_field("PUT")); ?>

            <?php echo e(csrf_field()); ?>

            <div class="panel">
                <?php $__currentLoopData = $settings; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $setting): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                    <div class="panel-heading">
                        <h3 class="panel-title">
                            <?php echo e($setting->display_name); ?><code>Voyager::setting('<?php echo e($setting->key); ?>')</code>
                        </h3>
                        <div class="panel-actions">
                            <a href="<?php echo e(route('voyager.settings.move_up', $setting->id)); ?>">
                                <i class="sort-icons voyager-sort-asc"></i>
                            </a>
                            <a href="<?php echo e(route('voyager.settings.move_down', $setting->id)); ?>">
                                <i class="sort-icons voyager-sort-desc"></i>
                            </a>
                            <i class="voyager-trash"
                               data-id="<?php echo e($setting->id); ?>"
                               data-display-key="<?php echo e($setting->key); ?>"
                               data-display-name="<?php echo e($setting->display_name); ?>"></i>
                        </div>
                    </div>
                    <div class="panel-body">
                        <?php if($setting->type == "text"): ?>
                            <input type="text" class="form-control" name="<?php echo e($setting->key); ?>" value="<?php echo e($setting->value); ?>">
                        <?php elseif($setting->type == "text_area"): ?>
                            <textarea class="form-control" name="<?php echo e($setting->key); ?>"><?php if(isset($setting->value)): ?><?php echo e($setting->value); ?><?php endif; ?></textarea>
                        <?php elseif($setting->type == "rich_text_box"): ?>
                            <textarea class="form-control richTextBox" name="<?php echo e($setting->key); ?>"><?php if(isset($setting->value)): ?><?php echo e($setting->value); ?><?php endif; ?></textarea>
                        <?php elseif($setting->type == "code_editor"): ?>
                            <?php $options = json_decode($setting->details); ?>
                            <div id="<?php echo e($setting->key); ?>" data-theme="<?php echo e(@$options->theme); ?>" data-language="<?php echo e(@$options->language); ?>" class="ace_editor min_height_400" name="<?php echo e($setting->key); ?>"><?php if(isset($setting->value)): ?><?php echo e($setting->value); ?><?php endif; ?></div>
                            <textarea name="<?php echo e($setting->key); ?>" id="<?php echo e($setting->key); ?>_textarea" class="hidden"><?php if(isset($setting->value)): ?><?php echo e($setting->value); ?><?php endif; ?></textarea>
                        <?php elseif($setting->type == "image" || $setting->type == "file"): ?>
                            <?php if(isset( $setting->value ) && !empty( $setting->value ) && Storage::disk(config('voyager.storage.disk'))->exists($setting->value)): ?>
                                <div class="img_settings_container">
                                    <a href="<?php echo e(route('voyager.settings.delete_value', $setting->id)); ?>" class="voyager-x"></a>
                                    <img src="<?php echo e(Storage::disk(config('voyager.storage.disk'))->url($setting->value)); ?>" style="width:200px; height:auto; padding:2px; border:1px solid #ddd; margin-bottom:10px;">
                                </div>
                                <div class="clearfix"></div>
                            <?php elseif($setting->type == "file" && isset( $setting->value )): ?>
                                <div class="fileType"><?php echo e($setting->value); ?></div>
                            <?php endif; ?>
                            <input type="file" name="<?php echo e($setting->key); ?>">
                        <?php elseif($setting->type == "select_dropdown"): ?>
                            <?php $options = json_decode($setting->details); ?>
                            <?php $selected_value = (isset($setting->value) && !empty($setting->value)) ? $setting->value : NULL; ?>
                            <select class="form-control" name="<?php echo e($setting->key); ?>">
                                <?php $default = (isset($options->default)) ? $options->default : NULL; ?>
                                <?php if(isset($options->options)): ?>
                                    <?php $__currentLoopData = $options->options; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $index => $option): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                        <option value="<?php echo e($index); ?>" <?php if($default == $index && $selected_value === NULL): ?><?php echo e('selected="selected"'); ?><?php endif; ?> <?php if($selected_value == $index): ?><?php echo e('selected="selected"'); ?><?php endif; ?>><?php echo e($option); ?></option>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                <?php endif; ?>
                            </select>

                        <?php elseif($setting->type == "radio_btn"): ?>
                            <?php $options = json_decode($setting->details); ?>
                            <?php $selected_value = (isset($setting->value) && !empty($setting->value)) ? $setting->value : NULL; ?>
                            <?php $default = (isset($options->default)) ? $options->default : NULL; ?>
                            <ul class="radio">
                                <?php if(isset($options->options)): ?>
                                    <?php $__currentLoopData = $options->options; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $index => $option): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                        <li>
                                            <input type="radio" id="option-<?php echo e($index); ?>" name="<?php echo e($setting->key); ?>"
                                                   value="<?php echo e($index); ?>" <?php if($default == $index && $selected_value === NULL): ?><?php echo e('checked'); ?><?php endif; ?> <?php if($selected_value == $index): ?><?php echo e('checked'); ?><?php endif; ?>>
                                            <label for="option-<?php echo e($index); ?>"><?php echo e($option); ?></label>
                                            <div class="check"></div>
                                        </li>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                <?php endif; ?>
                            </ul>
                        <?php elseif($setting->type == "checkbox"): ?>
                            <?php $options = json_decode($setting->details); ?>
                            <?php $checked = (isset($setting->value) && $setting->value == 1) ? true : false; ?>
                            <?php if(isset($options->on) && isset($options->off)): ?>
                                <input type="checkbox" name="<?php echo e($setting->key); ?>" class="toggleswitch" <?php if($checked): ?> checked <?php endif; ?> data-on="<?php echo e($options->on); ?>" data-off="<?php echo e($options->off); ?>">
                            <?php else: ?>
                                <input type="checkbox" name="<?php echo e($setting->key); ?>" <?php if($checked): ?> checked <?php endif; ?> class="toggleswitch">
                            <?php endif; ?>
                        <?php endif; ?>

                    </div>
                    <?php if(!$loop->last): ?>
                        <hr>
                    <?php endif; ?>
                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
            </div>
            <button type="submit" class="btn btn-primary pull-right">Save Settings</button>
        </form>

        <div style="clear:both"></div>

        <div class="panel" style="margin-top:10px;">
            <div class="panel-heading new-setting">
                <hr>
                <h3 class="panel-title"><i class="voyager-plus"></i> New Setting</h3>
            </div>
            <div class="panel-body">
                <form action="<?php echo e(route('voyager.settings.store')); ?>" method="POST">
                    <?php echo e(csrf_field()); ?>

                    <div class="col-md-4">
                        <label for="display_name">Name</label>
                        <input type="text" class="form-control" name="display_name" placeholder="Setting name ex: Admin Title" required="required">
                    </div>
                    <div class="col-md-4">
                        <label for="key">Key</label>
                        <input type="text" class="form-control" name="key" placeholder="Setting key ex: admin_title" required="required">
                    </div>
                    <div class="col-md-4">
                        <label for="asdf">Type</label>
                        <select name="type" class="form-control" required="required">
                            <option value="">Choose type</option>
                            <option value="text">Text Box</option>
                            <option value="text_area">Text Area</option>
                            <option value="rich_text_box">Rich Textbox</option>
                            <option value="code_editor">Code Editor</option>
                            <option value="checkbox">Check Box</option>
                            <option value="radio_btn">Radio Button</option>
                            <option value="select_dropdown">Select Dropdown</option>
                            <option value="file">File</option>
                            <option value="image">Image</option>
                        </select>
                    </div>
                    <div class="col-md-12">
                        <a id="toggle_options"><i class="voyager-double-down"></i> OPTIONS</a>
                        <div class="new-settings-options">
                            <label for="options">Options
                                <small>(optional, only applies to certain types like dropdown box or radio button)
                                </small>
                            </label>
                            <div id="options_editor" class="form-control min_height_200" data-language="json"></div>
                            <textarea id="options_textarea" name="details" class="hidden"></textarea>
                            <div id="valid_options" class="alert-success alert" style="display:none">Valid Json</div>
                            <div id="invalid_options" class="alert-danger alert" style="display:none">Invalid Json</div>
                        </div>
                    </div>
                    <script>
                        $('document').ready(function () {
                            $('#toggle_options').click(function () {
                                $('.new-settings-options').toggle();
                                if ($('#toggle_options .voyager-double-down').length) {
                                    $('#toggle_options .voyager-double-down').removeClass('voyager-double-down').addClass('voyager-double-up');
                                } else {
                                    $('#toggle_options .voyager-double-up').removeClass('voyager-double-up').addClass('voyager-double-down');
                                }
                            });
                        });
                    </script>
                    <div style="clear:both"></div>
                    <button type="submit" class="btn btn-primary pull-right new-setting-btn">
                        <i class="voyager-plus"></i> Add New Setting
                    </button>
                    <div style="clear:both"></div>
                </form>
            </div>
        </div>
    </div>

    <div class="modal modal-danger fade" tabindex="-1" id="delete_modal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    <h4 class="modal-title">
                        <i class="voyager-trash"></i> Are you sure you want to delete the <span id="delete_setting_title"></span> Setting?
                    </h4>
                </div>
                <div class="modal-footer">
                    <form action="<?php echo e(route('voyager.settings.delete', ['id' => '__id'])); ?>" id="delete_form" method="POST">
                        <?php echo e(method_field("DELETE")); ?>

                        <?php echo e(csrf_field()); ?>

                        <input type="submit" class="btn btn-danger pull-right delete-confirm" value="Yes, Delete This Setting">
                    </form>
                    <button type="button" class="btn btn-default pull-right" data-dismiss="modal">Cancel</button>
                </div>
            </div>
        </div>
    </div>

    <script>
        $('document').ready(function () {
            $('.voyager-trash').click(function () {
                var display = $(this).data('display-name') + '/' + $(this).data('display-key');

                $('#delete_setting_title').text(display);
                $('#delete_form')[0].action = $('#delete_form')[0].action.replace('__id', $(this).data('id'));
                $('#delete_modal').modal('show');
            });

            $('.toggleswitch').bootstrapToggle();
        });
    </script>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <iframe id="form_target" name="form_target" style="display:none"></iframe>
    <form id="my_form" action="<?php echo e(route('voyager.upload')); ?>" target="form_target" method="POST" enctype="multipart/form-data" style="width:0;height:0;overflow:hidden">
        <?php echo e(csrf_field()); ?>

        <input name="image" id="upload_file" type="file" onchange="$('#my_form').submit();this.value='';">
        <input type="hidden" name="type_slug" id="type_slug" value="settings">
    </form>

    <script src="<?php echo e(voyager_asset('lib/js/tinymce/tinymce.min.js')); ?>"></script>
    <script src="<?php echo e(voyager_asset('js/voyager_tinymce.js')); ?>"></script>
    <script src="<?php echo e(voyager_asset('lib/js/ace/ace.js')); ?>"></script>
    <script src="<?php echo e(voyager_asset('js/voyager_ace_editor.js')); ?>"></script>
    <script>
        var options_editor = ace.edit('options_editor');
        options_editor.getSession().setMode("ace/mode/json");

        var options_textarea = document.getElementById('options_textarea');
        options_editor.getSession().on('change', function() {
            options_textarea.value = options_editor.getValue();
        });
    </script>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>