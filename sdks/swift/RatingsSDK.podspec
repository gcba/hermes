Pod::Spec.new do |spec|
    spec.name = 'RatingsSDK'
    spec.version = '0.1.0'
    spec.summary = 'Hermes Swift SDK'
    spec.homepage = 'https://github.com/gcba/hermes'

    spec.authors = { 'Rita Zerrizuela' => 'zeta@widcket.com' }
    spec.license = { :type => 'MIT' }

    spec.ios.deployment_target = '8.0'
    spec.osx.deployment_target = '10.10'

    spec.source = { :http => 'https://github.com/gcba/hermes/raw/master/sdks/swift/RatingsSDK.zip' }

    spec.frameworks = 'Foundation'
    spec.dependency 'SwiftHTTP', '~> 2.0.2'
    spec.dependency 'SwifterSwift/Foundation', '~> 3.1.0'
    spec.dependency 'GBDeviceInfo', '~> 4.3.0'
end