import UIKit
import Alamofire

class CollectionViewDelegateAndDataSource: NSObject, UICollectionViewDelegate, UICollectionViewDataSource {

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
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "myCell", for: indexPath) as! CollectionViewCell
        cell.label.text = String(numbers[indexPath.row])
        return cell
    }

    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        if let cell = collectionView.cellForItem(at: indexPath) as? CollectionViewCell {
            cell.contentView.backgroundColor = #colorLiteral(red: 1, green: 0.4932718873, blue: 0.4739984274, alpha: 1)
            Alamofire.request("http://192.168.4.1/on/\(indexPath.row)")
        }
    }

    func collectionView(_ collectionView: UICollectionView, didDeselectItemAt indexPath: IndexPath) {
        // if let cell = collectionView.cellForItem(at: indexPath) as? CollectionViewCell {
        //     cell.contentView.backgroundColor = .white
        //     Alamofire.request("http://192.168.4.1/off/\(indexPath.row)")
        // }
    }

}
