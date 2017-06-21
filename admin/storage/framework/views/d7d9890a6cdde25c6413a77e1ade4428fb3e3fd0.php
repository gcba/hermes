<!DOCTYPE html>
<html>
<head>
    <title><?php echo $__env->yieldContent('page_title', setting('admin_title') . " - " . setting('admin_description')); ?></title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="csrf-token" content="<?php echo e(csrf_token()); ?>"/>

    <!-- Google Fonts -->
    <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,700" rel="stylesheet">

    <!-- Favicon -->
    <link rel="shortcut icon" href="<?php echo e(voyager_asset('images/logo-icon.png')); ?>" type="image/x-icon">

    

    <!-- App CSS -->
    <link rel="stylesheet" href="<?php echo e(voyager_asset('css/app.css')); ?>">

    <?php echo $__env->yieldContent('css'); ?>

    <!-- Few Dynamic Styles -->
    <style type="text/css">
        .voyager .side-menu .navbar-header {
            background:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
            border-color:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
        .widget .btn-primary{
            border-color:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
        .widget .btn-primary:focus, .widget .btn-primary:hover, .widget .btn-primary:active, .widget .btn-primary.active, .widget .btn-primary:active:focus{
            background:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
        .voyager .breadcrumb a{
            color:<?php echo e(config('voyager.primary_color','#22A7F0')); ?>;
        }
    </style>

    <?php if(!empty(config('voyager.additional_css'))): ?><!-- Additional CSS -->
        <?php $__currentLoopData = config('voyager.additional_css'); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $css): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?><link rel="stylesheet" type="text/css" href="<?php echo e(asset($css)); ?>"><?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
    <?php endif; ?>

    <?php echo $__env->yieldContent('head'); ?>
</head>

<body class="voyager <?php if(isset($dataType) && isset($dataType->slug)): ?><?php echo e($dataType->slug); ?><?php endif; ?>">

<div id="voyager-loader">
    <?php $admin_loader_img = Voyager::setting('admin_loader', ''); ?>
    <?php if($admin_loader_img == ''): ?>
        <img src="<?php echo e(voyager_asset('images/logo-icon.png')); ?>" alt="Voyager Loader">
    <?php else: ?>
        <img src="<?php echo e(Voyager::image($admin_loader_img)); ?>" alt="Voyager Loader">
    <?php endif; ?>
</div>

<?php
$user_avatar = Voyager::image(Auth::user()->avatar);
if ((substr(Auth::user()->avatar, 0, 7) == 'http://') || (substr(Auth::user()->avatar, 0, 8) == 'https://')) {
    $user_avatar = Auth::user()->avatar;
}
?>

<div class="app-container">
    <div class="fadetoblack visible-xs"></div>
    <div class="row content-container">
        <?php echo $__env->make('voyager::dashboard.navbar', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <?php echo $__env->make('voyager::dashboard.sidebar', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <!-- Main Content -->
        <div class="container-fluid">
            <div class="side-body padding-top">
                <?php echo $__env->yieldContent('page_header'); ?>
                <?php echo $__env->yieldContent('content'); ?>
            </div>
        </div>
    </div>
</div>
<?php echo $__env->make('voyager::partials.app-footer', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
<script>
    (function(){
            var appContainer = document.querySelector('.app-container'),
                sidebar = appContainer.querySelector('.side-menu'),
                navbar = appContainer.querySelector('nav.navbar.navbar-top'),
                loader = document.getElementById('voyager-loader'),
                anchor = document.getElementById('sidebar-anchor'),
                hamburgerMenu = document.querySelector('.hamburger'),
                sidebarTransition = sidebar.style.transition,
                navbarTransition = navbar.style.transition,
                containerTransition = appContainer.style.transition;

            sidebar.style.WebkitTransition = sidebar.style.MozTransition = sidebar.style.transition =
            appContainer.style.WebkitTransition = appContainer.style.MozTransition = appContainer.style.transition = 
            navbar.style.WebkitTransition = navbar.style.MozTransition = navbar.style.transition = 'none';
            
            if (window.localStorage && window.localStorage['voyager.stickySidebar'] == 'true') {
                appContainer.className += ' expanded';
                loader.style.left = (sidebar.clientWidth/2)+'px';
                anchor.className += ' active';
                anchor.dataset.sticky = anchor.title;
                anchor.title = anchor.dataset.unstick;
                hamburgerMenu.className += ' is-active';
            }

            navbar.style.WebkitTransition = navbar.style.MozTransition = navbar.style.transition = navbarTransition;
            sidebar.style.WebkitTransition = sidebar.style.MozTransition = sidebar.style.transition = sidebarTransition;
            appContainer.style.WebkitTransition = appContainer.style.MozTransition = appContainer.style.transition = containerTransition;
    })();
</script>
<!-- Javascript Libs -->


<script type="text/javascript" src="<?php echo e(voyager_asset('js/app.js')); ?>"></script>


<script>
    <?php if(Session::has('alerts')): ?>
        let alerts = <?php echo json_encode(Session::get('alerts')); ?>;
        displayAlerts(alerts, toastr);
    <?php endif; ?>

    <?php if(Session::has('message')): ?>
    
    // TODO: change Controllers to use AlertsMessages trait... then remove this
    var alertType = <?php echo json_encode(Session::get('alert-type', 'info')); ?>;
    var alertMessage = <?php echo json_encode(Session::get('message')); ?>;
    var alerter = toastr[alertType];

    if (alerter) {
        alerter(alertMessage);
    } else {
        toastr.error("toastr alert-type " + alertType + " is unknown");
    }

    <?php endif; ?>
</script>
<?php echo $__env->yieldContent('javascript'); ?>

<?php if(!empty(config('voyager.additional_js'))): ?><!-- Additional Javascript -->
    <?php $__currentLoopData = config('voyager.additional_js'); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $js): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?><script type="text/javascript" src="<?php echo e(asset($js)); ?>"></script><?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
<?php endif; ?>

</body>
</html>