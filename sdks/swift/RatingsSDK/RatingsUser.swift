//
//  RatingsUser.swift
//  RatingsSDK
//
//  Created by Rita Zerrizuela on 10/13/17.
//  Copyright © 2017 GCBA. All rights reserved.
//

import Foundation

public struct RatingsUser {
    public let name: String
    public let email: String?
    public let mibaId: String?
    
    public init(name: String, email: String? = nil, mibaId: String? = nil) {
        self.name = name
        self.email = email
        self.mibaId = mibaId
    }
    
    public func toDict() -> [String: String] {
        var user: [String: String] = [:]
        
        user["name"] = name
        
        if let mibaId = mibaId {
            user["mibaId"] = mibaId
        }
        
        if let email = email {
            user["email"] = email
        }
        
        return user
    }
}
