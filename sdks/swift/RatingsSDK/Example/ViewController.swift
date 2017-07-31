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
        // Do any additional setup after loading the view, typically from a nib.

        let app = "e10adc3949ba59abbe56e057f20f883e"
        let platform = app
        let range = app
        let token = app

        do {
            let sdk = try Rating(api: "https://7333ab98.ngrok.io", app: app, platform: platform, range: range, token: token)

            // Rating only; no user

            try sdk.create(rating: 5) { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }

            // Rating and description only; no user

            try sdk.create(rating: 4, description: "Bueno") { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }

            // Rating, description and comment; no user

            try sdk.create(rating: 3, description: "Regular", comment: "Lorem ipsum dolor...") { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }

            // Rating, description and comment; user name and mibaId only

            if let error = sdk.setUser(name: "Juan Pérez", mibaId: "e10adc3949") {
                debugPrint(error.localizedDescription)
                
                return
            }

            try sdk.create(rating: 2, description: "Malo", comment: "Lorem ipsum dolor...") { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }

            // Rating, description and comment; user name and email only
            
            if let error = sdk.setUser(name: "Juan Pérez", email: "juan@example.com") {
                debugPrint(error.localizedDescription)
                
                return
            }

            try sdk.create(rating: 1, description: "Muy Malo", comment: "Lorem ipsum dolor...") { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }

            // Rating, description and comment; user name, email and mibaId
            
            if let error = sdk.setUser(name: "Juan Pérez", mibaId: "e10adc3949", email: "juan@example.com") {
                debugPrint(error.localizedDescription)
                
                return
            }

            try sdk.create(rating: 5, description: "Muy Bueno", comment: "Lorem ipsum dolor...") { response in
                if let error = response.error {
                    debugPrint("Request error: \(error.localizedDescription)")
                }
            }
        } catch let error {
            debugPrint("Error: \(error.localizedDescription)")
        }
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
}

