//
//  ViewController.swift
//  Example
//
//  Created by Rita Zerrizuela on 7/25/17.
//  Copyright © 2017 GCBA. All rights reserved.
//

import UIKit
import RatingsSDK

class ViewController: UIViewController {
    override func viewDidLoad() {
        super.viewDidLoad()

        let app = "e10adc3949ba59abbe56e057f20f883e"
        let platform = app
        let range = app
        let token = app
        let sdk = Ratings(api: "https://73cbd954.ngrok.io", token: token, app: app, platform: platform, range: range)
        
        // Rating only; no user
        sdk.create(rating: 5) { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
        
        // Rating and description only; no user
        sdk.create(rating: 4, description: "Bueno") { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
        
        // Rating, description and comment; no user
        sdk.create(rating: 3, description: "Regular", comment: "Lorem ipsum dolor...") { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
        
        // Rating, description and comment; user name and mibaId only
        sdk.user = RatingsUser(name: "Juan Pérez", mibaId: "04860d65-7e93-49e8-a983-a4007d23ffa5")

        sdk.create(rating: 2, description: "Malo", comment: "Lorem ipsum dolor...") { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
        
        // Rating, description and comment; user name and email only
        sdk.user = RatingsUser(name: "Juan Pérez", email: "juan@example.com")
        
        sdk.create(rating: 1, description: "Muy Malo", comment: "Lorem ipsum dolor...") { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
        
        // Rating, description and comment; user name, email and mibaId
        sdk.user = RatingsUser(name: "Juan Pérez", email: "juan@example.com", mibaId: "08108a49-4c68-47da-8510-93922b6b2d76")
        
        sdk.create(rating: 5, description: "Muy Bueno", comment: "Lorem ipsum dolor...") { response, error in
            if error != nil {
                debugPrint("Error: \(error!.message)")
            }
        }
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
}

