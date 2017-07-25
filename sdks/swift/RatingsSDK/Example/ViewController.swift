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
        
        RatingsSDK.hello()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

