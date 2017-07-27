Pod::Spec.new do |spec|
    spec.name = 'RatingsSDK'
    spec.version = '1.0.0'
    spec.summary = 'Hermes Swift SDK'
    spec.homepage = 'https://github.com/gcba/hermes'

    spec.authors = { 'Rita Zerrizuela' => 'zeta@widcket.com' }
    spec.license = { :type => 'MIT', :file => 'LICENSE' }

    spec.ios.deployment_target = '8.0'
    spec.osx.deployment_target = '10.10'

    spec.source = { :git => 'https://github.com/gcba/hermes.git', :branch => 'master' }
    # spec.source_files = 'sdks/swift/RatingsSDK/RatingsSDK/*.{m,h,swift}'
    spec.vendored_frameworks = 'sdks/swift/RatingsSDK/RatingsSDK.framework'

    spec.frameworks = 'Foundation'
    spec.dependency 'SwiftHTTP', '~> 2.0.2'
    spec.dependency 'SwifterSwift/Foundation', '~> 3.1.0'
    spec.dependency 'GBDeviceInfo', '~> 4.3.0'
end