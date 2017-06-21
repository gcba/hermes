<!DOCTYPE html>
<html lang="<?php echo e(config('app.locale')); ?>">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="robots" content="none" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimal-ui">
    <meta name="description" content="admin login">
    <title>Admin - <?php echo e(Voyager::setting("title")); ?></title>
    <link rel="stylesheet" href="<?php echo e(voyager_asset('css/app.css')); ?>">
    <style>
        body {
            background-image:url('<?php echo e(Voyager::image( Voyager::setting("admin_bg_image"), config('voyager.assets_path') . "/images/bg.jpg" )); ?>');
            background-color: <?php echo e(Voyager::setting("admin_bg_color", "#FFFFFF" )); ?>;
        }
        .login-sidebar{
            border-top:5px solid <?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
        @media (max-width: 767px) {
            .login-sidebar {
                border-top:0px !important;
            }
        }
        .login-sidebar:after {
            background: linear-gradient(-135deg, <?php echo e(config('voyager.login.gradient_a','#ffffff')); ?>, <?php echo e(config('voyager.login.gradient_b','#ffffff')); ?>);
            background: -webkit-linear-gradient(-135deg, <?php echo e(config('voyager.login.gradient_a','#ffffff')); ?>, <?php echo e(config('voyager.login.gradient_b','#ffffff')); ?>);
        }
        .login-button, .bar:before, .bar:after{
            background:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
    </style>

    <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,700" rel="stylesheet">
</head>
<body class="login">
<div class="container-fluid">
    <div class="row">
        <div class="faded-bg animated"></div>
        <div class="hidden-xs col-sm-7 col-md-8">
            <div class="clearfix">
                <div class="col-sm-12 col-md-10 col-md-offset-2">
                    <div class="logo-title-container">
                        <?php $admin_logo_img = Voyager::setting('admin_icon_image', ''); ?>
                        <?php if($admin_logo_img == ''): ?>
                        <img class="img-responsive pull-left logo hidden-xs animated fadeIn" src="<?php echo e(voyager_asset('images/logo-icon-light.png')); ?>" alt="Logo Icon">
                        <?php else: ?>
                        <img class="img-responsive pull-left logo hidden-xs animated fadeIn" src="<?php echo e(Voyager::image($admin_logo_img)); ?>" alt="Logo Icon">
                        <?php endif; ?>
                        <div class="copy animated fadeIn">
                            <h1><?php echo e(Voyager::setting('admin_title', 'Voyager')); ?></h1>
                            <p><?php echo e(Voyager::setting('admin_description', __('voyager.login.welcome'))); ?></p>
                        </div>
                    </div> <!-- .logo-title-container -->
                </div>
            </div>
        </div>

        <div class="col-xs-12 col-sm-5 col-md-4 login-sidebar">

            <div class="login-container">
                <h2><?php echo e(__('voyager.login.signin_below')); ?></h2>
                <form action="<?php echo e(route('voyager.login')); ?>" method="POST">
                <?php echo e(csrf_field()); ?>

                <div class="group">
                  <input type="text" name="email" value="<?php echo e(old('email')); ?>" required>
                  <span class="highlight"></span>
                  <span class="bar"></span>
                  <label><i class="glyphicon glyphicon-user"></i><span class="span-input"> <?php echo e(__('voyager.generic.email')); ?></span></label>
                </div>

                <div class="group">
                  <input type="password" name="password" required>
                  <span class="highlight"></span>
                  <span class="bar"></span>
                  <label><i class="glyphicon glyphicon-lock"></i><span class="span-input"> <?php echo e(__('voyager.generic.password')); ?></span></label>
                </div>
                
                    <button type="submit" class="btn btn-block login-button">
                        <span class="signingin hidden"><span class="voyager-refresh"></span> <?php echo e(__('voyager.login.logginin')); ?>...</span>
                        <span class="signin"><?php echo e(__('voyager.generic.login')); ?></span>
                    </button>
                    
              </form>

              <div style="clear:both"></div>

              <?php if(!$errors->isEmpty()): ?>
              <div class="alert alert-red">
                <ul class="list-unstyled">
                    <?php $__currentLoopData = $errors->all(); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $err): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                    <li><?php echo e($err); ?></li>
                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                </ul>
              </div>
              <?php endif; ?>

            </div> <!-- .login-container -->

        </div> <!-- .login-sidebar -->
    </div> <!-- .row -->
</div> <!-- .container-fluid -->
<script>
    var btn = document.querySelector('button[type="submit"]');
    var form = document.forms[0];
    btn.addEventListener('click', function(ev){
        if (form.checkValidity()) {
            btn.querySelector('.signingin').className = 'signingin';
            btn.querySelector('.signin').className = 'signin hidden';
        } else {
            ev.preventDefault();
        }
    });
    document.getElementById('email').focus();
</script>
</body>
</html>
