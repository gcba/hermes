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
    case invalidName(String)
    case invalidEmail(String)
    case invalidMibaId(String)
    case invalidDescription(String)
    case invalidRating(String)
    case invalidComment(String)
    case missingEmailAndMibaId(String)
    case http(String, Int?, Response?)
    case other(String)
}

public extension RatingsError {
    public var message: String {
        switch self {
        case .invalidName(let description): return description
        case .invalidEmail(let description): return description
        case .invalidMibaId(let description): return description
        case .invalidDescription(let description): return description
        case .invalidRating(let description): return description
        case .invalidComment(let description): return description
        case .missingEmailAndMibaId(let description): return description
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
