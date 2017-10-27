//
//  RatingsError.swift
//  RatingsSDK
//
//  Created by Rita Zerrizuela on 10/13/17.
//  Copyright © 2017 GCBA. All rights reserved.
//

import Foundation
import SwiftHTTP

public enum RatingsError: Error {
    case nameTooShort(String)
    case nameTooLong(String)
    case invalidEmail(String)
    case emailTooShort(String)
    case emailTooLong(String)
    case invalidMibaId(String)
    case missingEmailAndMibaId(String)
    case descriptionTooShort(String)
    case descriptionTooLong(String)
    case commentTooShort(String)
    case commentTooLong(String)
    case invalidRating(String)
    case http(String, Int?, Response?)
    case other(String)
}

public extension RatingsError {
    public var message: String {
        switch self {
        case .nameTooShort(let description): return description
        case .nameTooLong(let description): return description
        case .invalidEmail(let description): return description
        case .emailTooShort(let description): return description
        case .emailTooLong(let description): return description
        case .invalidMibaId(let description): return description
        case .missingEmailAndMibaId(let description): return description
        case .descriptionTooShort(let description): return description
        case .descriptionTooLong(let description): return description
        case .commentTooShort(let description): return description
        case .commentTooLong(let description): return description
        case .invalidRating(let description): return description
        case .http(let description, _, _): return description
        case .other(let description): return description
        }
    }
    
    public var statusCode: Int? {
        switch self {
        case .http(_, let status, _): return status
        default: return nil
        }
    }
    
    public var response: Response? {
        switch self {
        case .http(_, _, let res): return res
        default: return nil
        }
    }
}
