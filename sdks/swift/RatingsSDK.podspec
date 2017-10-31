Pod::Spec.new do |spec|
    spec.name = 'RatingsSDK'
    spec.version = '0.1.0'
    spec.summary = 'Hermes Swift SDK'
    spec.homepage = 'https://gcba.github.io/hermes/uso/sdks.html'
    spec.authors = { 'Rita Zerrizuela' => 'zeta@widcket.com' }
    spec.license = { :type => 'MIT', :file => 'LICENSE.md' }
    spec.source = { :git => 'https://github.com/gcba/hermes.git', :branch => 'master' }
    spec.source_files = 'sdks/swift/RatingsSDK/Ratings*.swift'
    spec.frameworks = 'Foundation'
    spec.vendored_frameworks = ['sdks/swift/Rome/*.framework']
    spec.ios.deployment_target = '9.0'
end