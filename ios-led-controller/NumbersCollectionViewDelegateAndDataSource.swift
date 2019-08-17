import UIKit
import Alamofire

class NumbersCollectionViewDelegateAndDataSource: NSObject, UICollectionViewDelegate, UICollectionViewDataSource {

    let numbers: [Int]
    
    override init() {
        var nums: [Int] = []
        for i in 0...760 {
            nums.append(i)
        }
        self.numbers = nums
    }
    
    func numberOfSections(in collectionView: UICollectionView) -> Int {
        return 1
    }
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return numbers.count
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "myCell", for: indexPath) as! NumbersCollectionViewCell
        cell.label.text = String(numbers[indexPath.row])
        return cell
    }

    // func collectionView(_ collectionView: UICollectionView, shouldHighlightItemAt indexPath: IndexPath) -> Bool {
    //     return true
    // }

    // func collectionView(_ collectionView: UICollectionView, didHighlightItemAt indexPath: IndexPath) {
    //     Alamofire.request("http://192.168.4.1/\(indexPath.row)")
    // }

    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        // let url = thumbnailFileURLS[indexPath.item]
        // if UIApplication.sharedApplication().canOpenURL(url) {
        //     UIApplication.sharedApplication().openURL(url)
        // }
        Alamofire.request("http://192.168.4.1/\(indexPath.row)")
    }


    
}
