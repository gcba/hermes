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

public class Ratings {
    public init(api url: String, app: String, platform: String, range: String, token: String) {
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
    
    private let url: String
    private let app: String
    private let platform: String
    private let range: String
    private let token: String
    private let deviceInfo: GBDeviceInfo
    
    private var timeout: Double
    private var _user: RatingUser?
    
    // MARK: - Public properties
    
    public var user: RatingUser?

    // MARK: - Validations

    private func validateUrl(_ url: String) {
        guard url.isValidUrl else { fatalError("invalid url") }
    }
    
    private func validateKey(_ key: String, description: String) {
        guard key.length == 32 else { fatalError("\(description) is not a valid key") }
    }
    
    private func validateToken(_ token: String) {
        guard token.length > 0 else { fatalError("invalid token") }
    }
    
    private func validateName(_ name: String) -> RatingError? {
        guard name.length >= 3 else { return RatingError.invalidName("name too short") }
        guard name.length <= 70 else { return RatingError.invalidName("name too long") }
        
        return nil
    }
    
    private func validateEmail(_ email: String) -> RatingError? {
        guard email.isEmail else { return RatingError.invalidEmail("invalid email") }
        guard email.length >= 3 else { return RatingError.invalidEmail("email too short") }
        guard email.length <= 100 else { return RatingError.invalidEmail("email too long") }
        
        return nil
    }
    
    private func validateMibaId(_ mibaId: String) -> RatingError? {
        guard mibaId.length == 36 else { return RatingError.invalidMibaId("invalid mibaId") }
        
        return nil
    }
    
    private func validateRating(_ rating: Int) -> RatingError? {
        guard rating >= -127 && rating <= 127 else { return RatingError.invalidRating("invalid rating") }
        
        return nil
    }
    
    private func validateDescription(_ description: String) -> RatingError? {
        guard description.length >= 3 else { return RatingError.invalidDescription("description too short") }
        guard description.length <= 30 else { return RatingError.invalidDescription("description too long") }
        
        return nil
    }
    
    private func validateComment(_ comment: String) -> RatingError? {
        guard comment.length >= 3 else { return RatingError.invalidComment("comment too short") }
        guard comment.length <= 1000 else { return RatingError.invalidComment("comment too long") }
        
        return nil
    }

    // MARK: - Setters
    
    public func validateUser(_ user: RatingUser) -> RatingError? {
        guard user.mibaId != nil || user.email != nil else { return RatingError.missingEmailAndMibaId("user is missing email and mibaId") }
        
        if let error = validateName(user.name) { return error }
        
        if let mibaId = user.mibaId {
            if let error = validateMibaId(mibaId) { return error }
        }
        
        if let email = user.email {
            if let error = validateEmail(email) { return error }
        }
        
        _user = user
        
        return nil
    }
    
    // MARK: - Helpers
    
    private func buildParams() -> [String: Any] {
        let screen: [String: Any] = [
            "width": UIScreen.main.bounds.width,
            "height": UIScreen.main.bounds.height,
            "ppi": deviceInfo.displayInfo.pixelsPerInch
        ]
        
        let device: [String: Any] = [
            "name": deviceInfo.modelString,
            "brand": "Apple",
            "screen": screen
        ]
        
        let result: [String: Any]  = [
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
    
    private func send(params: [String: Any], retry: Int = 3, callback: @escaping (_ response: Response?, _ error: RatingError?)->()) {
        let headers = [
            "Content-Type": "application/json; charset=UTF-8",
            "Accept": "application/json",
            "Authorization": "Bearer \(token)"
        ]
        
        let backoff = 3 - retry
        
        HTTP.globalRequest { [unowned self] (request: NSMutableURLRequest) in
            request.timeoutInterval = self.timeout
        }
        
        DispatchQueue.global(qos: .userInitiated).asyncAfter(deadline: .now() + .seconds(backoff)) {
            do {
                try HTTP.POST(self.url, parameters: params, headers: headers, requestSerializer: JSONParameterSerializer()).start { response in
                    if response.error != nil && (response.error!.code == 503 || response.error!.code == 504 || response.error!.code >= 520) {
                        guard retry > 0 else {
                            callback(response, RatingError.http("Could not create new rating", response.error!.code, response))
                            
                            return
                        }
                        
                        self.send(params: params, retry: retry - 1, callback: callback)
                    } else {
                        callback(response, nil)
                    }
                }
            }
            catch let error {
                debugPrint(error.localizedDescription)
                callback(nil, RatingError.other(error.localizedDescription))
            }
        }
    }
    
    // MARK: - Public API
    
    public func create(rating: Int, description: String? = nil, comment: String? = nil, callback: @escaping (_ response: Response?, _ error: RatingError?)->()) {
        if let error = validateRating(rating) { return callback(nil, error) }
        
        var params: [String: Any] = buildParams()
        
        params["rating"] = rating
        
        if let description = description {
            if let error = validateDescription(description) { callback(nil, error) }
            
            params["description"] = description
        }
        
        if let comment = comment {
            if let error = validateComment(comment) { callback(nil, error) }
            
            params["comment"] = comment
        }
        
        if let user = user {
            if let error = validateUser(user) { callback(nil, error) }
            
            params["user"] = user.toDict()
        }
        
        send(params: params, callback: callback)
    }
}
