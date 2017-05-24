<footer class="app-footer">
    <div class="site-footer-right">
        <?php if(rand(1,100) == 100): ?>
            <i class="voyager-rum-1"></i> Made with rum and even more rum
        <?php else: ?>
            Made with <i class="voyager-heart"></i> by <a href="http://thecontrolgroup.com" target="_blank">The Control Group</a>
        <?php endif; ?>
        <?php  $version = Voyager::getVersion();  ?>
        <?php if(!empty($version)): ?>
            - <?php echo e($version); ?>

        <?php endif; ?>
    </div>
</footer>
