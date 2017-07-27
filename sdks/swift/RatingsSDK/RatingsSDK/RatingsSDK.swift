//
//  RatingsSDK.swift
//  RatingsSDK
//
//  Created by Rita Zerrizuela on 7/25/17.
//  Copyright © 2017 GCBA. All rights reserved.
//

import Foundation
import SwiftHTTP
import SwifterSwift
import GBDeviceInfo

public enum RatingError: Error {
    case validation(message: String)
}

public class Rating {
    public init(api url: String, app: String, platform: String, range: String, token: String) throws {
        let baseUrl = url.trimmed
        
        self.url = baseUrl.lastCharacterAsString == "/" ? baseUrl + "ratings" :  baseUrl + "/ratings"
        self.app = app.trimmed
        self.platform = platform.trimmed
        self.range = range.trimmed
        self.token = token.trimmed
        self.deviceInfo = GBDeviceInfo()
        self.timeout = 3
        
        try validateUrl(self.url)
        try validateKey(self.app, description: "app")
        try validateKey(self.platform, description: "platform")
        try validateKey(self.range, description: "range")
        try validateToken(self.token)
    }
    
    // MARK: - Private properties
    
    let url: String
    let app: String
    let platform: String
    let range: String
    let token: String
    let deviceInfo: GBDeviceInfo
    
    var timeout: Double
    var user: [String: String]?

    // MARK: - Validations

    func validateUrl(_ url: String) throws {
        guard url.isValidUrl else {
             throw RatingError.validation(message: "invalid url")
        }
    }
    
    func validateKey(_ key: String, description: String) throws {
        guard key.length == 32 else {
            throw RatingError.validation(message: "\(description) is not a valid key")
        }
    }
    
    func validateToken(_ token: String) throws {
        guard token.length > 0 else {
            throw RatingError.validation(message: "invalid token")
        }
    }
    
    func validateName(_ name: String) throws {
        guard name.length >= 3 else {
            throw RatingError.validation(message: "name too short")
        }
        
        guard name.length <= 70 else {
            throw RatingError.validation(message: "name too long")
        }
    }
    
    func validateEmail(_ email: String) throws {
        guard email.isEmail else {
            throw RatingError.validation(message: "invalid email")
        }
        
        guard email.length >= 3 else {
            throw RatingError.validation(message: "email too short")
        }
        
        guard email.length <= 100 else {
            throw RatingError.validation(message: "email too long")
        }
    }
    
    func validateMibaId(_ mibaId: String) throws {
        guard mibaId.length > 0 else {
            throw RatingError.validation(message: "mibaId too short")
        }
    }
    
    func validateRating(_ rating: Int) throws {
        guard rating >= -127 && rating <= 127 else {
            throw RatingError.validation(message: "invalid rating")
        }
    }
    
    func validateDescription(_ description: String) throws {
        guard description.length >= 3 else {
            throw RatingError.validation(message: "description too short")
        }
        
        guard description.length <= 30 else {
            throw RatingError.validation(message: "description too long")
        }
    }
    
    func validateComment(_ comment: String) throws {
        guard comment.length >= 3 else {
            throw RatingError.validation(message: "comment too short")
        }
        
        guard comment.length <= 1000 else {
            throw RatingError.validation(message: "comment too long")
        }
    }

    // MARK: - Setters
    
    public func setUser(name: String?, mibaId: String? = nil, email: String? = nil) throws {
        guard mibaId != nil || email != nil else {
            throw RatingError.validation(message: "user has no valid email or mibaId")
        }
        
        var newUser: [String: String] = [:]
        
        if let actualName = name {
            try validateName(actualName)
            
            newUser["name"] = actualName
        }
        
        if let actualMibaId = mibaId {
            try validateMibaId(actualMibaId)
            
            newUser["mibaId"] = actualMibaId
        }
        
        if let actualEmail = email {
            try validateEmail(actualEmail)
            
            newUser["email"] = actualEmail
        }
        
        user = newUser
    }
    
    // MARK: - Helpers
    
    func buildParams() -> [String: Any] {
        let screen: [String: Any] = [
            "width": UIScreen.main.bounds.width,
            "height": UIScreen.main.bounds.height,
            "ppi": deviceInfo.displayInfo.pixelsPerInch
        ]
        
        let device: [String : Any] = [
            "name": deviceInfo.modelString,
            "brand": "Apple",
            "screen": screen
        ]
        
        let result: [String : Any]  = [
            "range": range,
            "app": [
                "key": app,
                "version": Bundle.main.object(forInfoDictionaryKey: "CFBundleShortVersionString") as! String
            ],
            "platform": [
                "key": platform,
                "version": "\(deviceInfo.osVersion.major).\(deviceInfo.osVersion.minor).\(deviceInfo.osVersion.patch)"
            ],
            "device": device
        ]
        
        return result
    }
    
    func send(params: [String: Any], callback: @escaping (_ response: Response)->()) throws {
        let headers = [
            "Content-Type": "application/json; charset=UTF-8",
            "Accept": "application/json",
            "Authorization": "Bearer \(token)"
        ]
        
        HTTP.globalRequest { [unowned self] (request: NSMutableURLRequest) in
            request.timeoutInterval = self.timeout
        }
        
        try HTTP.POST(self.url, parameters: params, headers: headers, requestSerializer: JSONParameterSerializer()).start { response in
            callback(response)
        }
    }
    
    // MARK: - Public API
    
    public func create(rating: Int, description: String? = nil, comment: String? = nil, callback: @escaping (_ response: Response)->()) throws {
        try validateRating(rating)
        
        var params: [String: Any] = buildParams()
        
        params["rating"] = rating
        
        if let actualDescription = description {
            try validateDescription(actualDescription)
            
            params["description"] = actualDescription
        }
        
        if let actualComment = comment {
            try validateComment(actualComment)
            
            params["comment"] = actualComment
        }
        
        if let actualUser = user {
            params["user"] = actualUser
        }
        
        return try send(params: params, callback: callback)
    }
}
