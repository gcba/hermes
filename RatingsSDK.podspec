Pod::Spec.new do |spec|
    spec.name = 'RatingsSDK'
    spec.version = '0.1.0'
    spec.summary = 'Hermes Swift SDK'
    spec.homepage = 'https://gcba.github.io/hermes/uso/sdks.html'
    spec.authors = { 'Rita Zerrizuela' => 'zeta@widcket.com' }
    spec.license = { :type => 'MIT', :file => 'LICENSE' }
    spec.platform = :ios, '9.0'
    spec.source = { :git => 'https://github.com/gcba/hermes.git', :branch => 'master' }
    spec.source_files = 'sdks/swift/RatingsSDK/Ratings*.{swift,h,m}'
    spec.public_header_files = 'sdks/swift/RatingsSDK/Ratings*.h'
    spec.frameworks = 'Foundation'

    spec.dependency 'SwiftHTTP', '~> 2.0.2'
    spec.dependency 'SwifterSwift/Foundation', '~> 3.1.0'
    spec.dependency 'GBDeviceInfo', '~> 4.3.0'
end