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
        
        validateUrl(self.url)
        validateKey(self.app, description: "app")
        validateKey(self.platform, description: "platform")
        validateKey(self.range, description: "range")
        validateToken(self.token)
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

    func validateUrl(_ url: String) {
        guard url.isValidUrl else {
             fatalError("invalid url")
        }
    }
    
    func validateKey(_ key: String, description: String) {
        guard key.length == 32 else {
            fatalError("\(description) is not a valid key")
        }
    }
    
    func validateToken(_ token: String) {
        guard token.length > 0 else {
            fatalError("invalid token")
        }
    }
    
    func validateName(_ name: String) -> RatingError? {
        guard name.length >= 3 else {
            return RatingError.validation(message: "name too short")
        }
        
        guard name.length <= 70 else {
            return RatingError.validation(message: "name too long")
        }
        
        return nil
    }
    
    func validateEmail(_ email: String) -> RatingError? {
        guard email.isEmail else {
            return RatingError.validation(message: "invalid email")
        }
        
        guard email.length >= 3 else {
            return RatingError.validation(message: "email too short")
        }
        
        guard email.length <= 100 else {
            return RatingError.validation(message: "email too long")
        }
        
        return nil
    }
    
    func validateMibaId(_ mibaId: String) -> RatingError? {
        guard mibaId.length > 0 else {
            return RatingError.validation(message: "mibaId too short")
        }
        
        return nil
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
    
    public func setUser(name: String?, mibaId: String? = nil, email: String? = nil) -> RatingError? {
        guard mibaId != nil || email != nil else {
            return RatingError.validation(message: "user has no valid email or mibaId")
        }
        
        var newUser: [String: String] = [:]
        
        if let actualName = name {
            if let error = validateName(actualName) {
                return error
            }
            
            newUser["name"] = actualName
        }
        
        if let actualMibaId = mibaId {
            if let error = validateMibaId(actualMibaId) {
                return error
            }
            
            newUser["mibaId"] = actualMibaId
        }
        
        if let actualEmail = email {
            if let error = validateEmail(actualEmail) {
                return error
            }

            newUser["email"] = actualEmail
        }
        
        user = newUser
        
        return nil
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
    
    func send(params: [String: Any], retry: Int = 3, callback: @escaping (_ response: Response)->()) {
        let headers = [
            "Content-Type": "application/json; charset=UTF-8",
            "Accept": "application/json",
            "Authorization": "Bearer \(token)"
        ]
        
        HTTP.globalRequest { [unowned self] (request: NSMutableURLRequest) in
            request.timeoutInterval = self.timeout
        }
        
        do {
            try HTTP.POST(self.url, parameters: params, headers: headers, requestSerializer: JSONParameterSerializer()).start { response in
                if response.error != nil {
                    guard retry > 0 else {
                        callback(response)
                        
                        return
                    }

                    self.send(params: params, retry: retry - 1, callback: callback)
                } else {
                    callback(response)
                }
            }
        }
        catch let error {
            debugPrint(error.localizedDescription)
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
        
        return send(params: params, callback: callback)
    }
}
