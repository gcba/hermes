Pod::Spec.new do |spec|
    spec.name = 'RatingsSDK'
    spec.version = '1.0.0'
    spec.summary = 'Hermes Swift SDK'
    spec.homepage = 'https://github.com/gcba/hermes'

    spec.authors = { 'Rita Zerrizuela' => 'zeta@widcket.com' }
    spec.license = { :type => 'MIT' }

    spec.platform = { :ios => "8.0", :osx => "10.7", :watchos => "2.0", :tvos => "9.0" }
    spec.source = { :http => 'https://github.com/gcba/hermes/raw/master/sdks/swift/RatingsSDK/RatingsSDK.zip' }

    spec.ios.vendored_frameworks = 'RatingsSDK.framework'
    spec.osx.vendored_frameworks = 'RatingsSDK.framework'
    spec.watchos.vendored_frameworks = 'RatingsSDK.framework'
    spec.tvos.vendored_frameworks = 'RatingsSDK.framework'
end