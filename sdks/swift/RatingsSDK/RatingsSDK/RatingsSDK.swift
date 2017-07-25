//
//  RatingsSDK.swift
//  RatingsSDK
//
//  Created by Rita Zerrizuela on 7/25/17.
//  Copyright © 2017 GCBA. All rights reserved.
//

import Foundation
import SwifterSwift

public class RatingsSDK {
    public class func hello(){
        debugPrint("Hello from RatingsSDK!")
    }
}

enum RatingError: Error {
    case validationError(message: String)
}

public class Rating {
    init(api url: String, app: String, platform: String, range: String, token: String) {
        self.url = url.trimmed
        self.app = app.trimmed
        self.platform = platform.trimmed
        self.range = range.trimmed
        self.token = token.trimmed
        
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
    
    var _user: [String: String]?
    
    // MARK: - Public properties
    
    public var user: [String: String]? {
        get {
            return _user
        }
    }
    
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
    
    func validateName(_ name: String) throws {
        guard name.length >= 3 else {
            throw RatingError.validationError(message: "name too short")
        }
        
        guard name.length <= 70 else {
            throw RatingError.validationError(message: "name too long")
        }
    }
    
    func validateEmail(_ email: String) throws {
        guard email.isEmail else {
            throw RatingError.validationError(message: "invalid email")
        }
        
        guard email.length >= 3 else {
            throw RatingError.validationError(message: "email too short")
        }
        
        guard email.length <= 100 else {
            throw RatingError.validationError(message: "email too long")
        }
    }
    
    func validateMibaId(_ mibaId: String) throws {
        guard mibaId.length > 0 else {
            throw RatingError.validationError(message: "mibaId too short")
        }
    }
    
    func validateRating(_ rating: Int) throws {
        return
    }
    
    func validateDescription() -> String? {
        return ""
    }
    
    func validateComment() -> String? {
        return ""
    }
    
    // MARK: - Setters
    
    func setUser(mibaId: String) throws {
        try validateMibaId(mibaId)
        
        _user = ["mibaId": mibaId]
    }
    
    func setUser(email: String) throws {
        try validateEmail(email)
        
        _user = ["email": email]
    }
    
    func setUser(name: String, mibaId: String) throws {
        try validateName(name)
        try validateMibaId(mibaId)
        
        _user = [
            "name": name,
            "mibaId": mibaId
        ]
    }
    
    func setUser(name: String, email: String) throws {
        try validateName(name)
        try validateEmail(email)
        
        _user = [
            "name": name,
            "email": email
        ]
    }
    
    func setUser(name: String, mibaId: String, email: String) throws {
        try validateName(name)
        try validateMibaId(mibaId)
        try validateEmail(email)
        
        _user = [
            "name": name,
            "mibaId": mibaId,
            "email": email
        ]
    }
}
