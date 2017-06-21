<?php $__env->startSection('css'); ?>
    <meta name="csrf-token" content="<?php echo e(csrf_token()); ?>">
<?php $__env->stopSection(); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="<?php echo e($dataType->icon); ?>"></i> <?php if(isset($dataTypeContent->id)): ?><?php echo e(__('voyager.generic.edit')); ?><?php else: ?><?php echo e(__('voyager.generic.new')); ?><?php endif; ?> <?php echo e($dataType->display_name_singular); ?>

    </h1>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>
    <div class="page-content container-fluid">
        <div class="row">
            <div class="col-md-12">

                <div class="panel panel-bordered">

                    <div class="panel-heading">
                        <h3 class="panel-title"><?php if(isset($dataTypeContent->id)): ?><?php echo e(__('voyager.generic.edit')); ?><?php else: ?><?php echo e(__('voyager.generic.add_new')); ?><?php endif; ?> <?php echo e($dataType->display_name_singular); ?></h3>
                    </div>
                    <!-- /.box-header -->
                    <!-- form start -->
                    <form class="form-edit-add" role="form"
                          action="<?php if(isset($dataTypeContent->id)): ?><?php echo e(route('voyager.'.$dataType->slug.'.update', $dataTypeContent->id)); ?><?php else: ?><?php echo e(route('voyager.'.$dataType->slug.'.store')); ?><?php endif; ?>"
                          method="POST" enctype="multipart/form-data">
                        <!-- PUT Method if we are editing -->
                        <?php if(isset($dataTypeContent->id)): ?>
                            <?php echo e(method_field("PUT")); ?>

                        <?php endif; ?>

                        <!-- CSRF TOKEN -->
                        <?php echo e(csrf_field()); ?>


                        <div class="panel-body">
                            <div class="form-group">
                                <label for="name"><?php echo e(__('voyager.generic.name')); ?></label>
                                <input type="text" class="form-control" name="name"
                                    placeholder="<?php echo e(__('voyager.generic.name')); ?>" id="name"
                                    value="<?php if(isset($dataTypeContent->name)): ?><?php echo e(old('name', $dataTypeContent->name)); ?><?php else: ?><?php echo e(old('name')); ?><?php endif; ?>">
                            </div>

                            <div class="form-group">
                                <label for="name"><?php echo e(__('voyager.generic.email')); ?></label>
                                <input type="text" class="form-control" name="email"
                                       placeholder="<?php echo e(__('voyager.generic.email')); ?>" id="email"
                                       value="<?php if(isset($dataTypeContent->email)): ?><?php echo e(old('email', $dataTypeContent->email)); ?><?php else: ?><?php echo e(old('email')); ?><?php endif; ?>">
                            </div>

                            <div class="form-group">
                                <label for="password"><?php echo e(__('voyager.profile.password')); ?></label>
                                <?php if(isset($dataTypeContent->password)): ?>
                                    <br>
                                    <small><?php echo e(__('voyager.profile.password_hint')); ?></small>
                                <?php endif; ?>
                                <input type="password" class="form-control" name="password"
                                       placeholder="<?php echo e(__('voyager.profile.password')); ?>" id="password"
                                       value="">
                            </div>

                            <div class="form-group">
                                <label for="password"><?php echo e(__('voyager.profile.avatar')); ?></label>
                                <?php if(isset($dataTypeContent->avatar)): ?>
                                    <img src="<?php echo e(Voyager::image( $dataTypeContent->avatar )); ?>"
                                         style="width:200px; height:auto; clear:both; display:block; padding:2px; border:1px solid #ddd; margin-bottom:10px;">
                                <?php endif; ?>
                                <input type="file" name="avatar">
                            </div>

                            <div class="form-group">
                                <label for="role"><?php echo e(__('voyager.profile.user_role')); ?></label>
                                <select name="role_id" id="role" class="form-control">
                                    <?php $roles = TCG\Voyager\Models\Role::all(); ?>
                                    <?php $__currentLoopData = $roles; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $role): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                        <option value="<?php echo e($role->id); ?>" <?php if(isset($dataTypeContent) && $dataTypeContent->role_id == $role->id): ?> selected <?php endif; ?>><?php echo e($role->display_name); ?></option>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                </select>
                            </div>



                        </div><!-- panel-body -->

                        <div class="panel-footer">
                            <button type="submit" class="btn btn-primary"><?php echo e(__('voyager.generic.submit')); ?></button>
                        </div>
                    </form>

                    <iframe id="form_target" name="form_target" style="display:none"></iframe>
                    <form id="my_form" action="<?php echo e(route('voyager.upload')); ?>" target="form_target" method="post"
                          enctype="multipart/form-data" style="width:0;height:0;overflow:hidden">
                        <input name="image" id="upload_file" type="file"
                               onchange="$('#my_form').submit();this.value='';">
                        <input type="hidden" name="type_slug" id="type_slug" value="<?php echo e($dataType->slug); ?>">
                        <?php echo e(csrf_field()); ?>

                    </form>

                </div>
            </div>
        </div>
    </div>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <script>
        $('document').ready(function () {
            $('.toggleswitch').bootstrapToggle();
        });
    </script>
    <script src="<?php echo e(voyager_asset('lib/js/tinymce/tinymce.min.js')); ?>"></script>
    <script src="<?php echo e(voyager_asset('js/voyager_tinymce.js')); ?>"></script>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>