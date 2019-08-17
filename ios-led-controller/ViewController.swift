import UIKit

class ViewController: UIViewController {

    var headerView: UIView!
    var titleLabel: UILabel!
    var collectionView: UICollectionView!
    let collectionViewDelegateAndDataSource = CollectionViewDelegateAndDataSource()

    override func viewDidLoad() {
        super.viewDidLoad()
        // setupHeaderAndTitleLabel()
        
        let frame = self.view.frame
        let layout = UICollectionViewFlowLayout()
        collectionView = UICollectionView(frame: frame, collectionViewLayout: layout)
        self.view.addSubview(collectionView)
        
        collectionView.translatesAutoresizingMaskIntoConstraints = false
        collectionView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor).isActive = true
        collectionView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor).isActive = true
        collectionView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor).isActive = true
        // collectionView.topAnchor.constraint(equalTo: headerView.bottomAnchor).isActive = true
        collectionView.topAnchor.constraint(equalTo: self.view.topAnchor).isActive = true
    
    
        collectionView.register(CollectionViewCell.self, forCellWithReuseIdentifier: "myCell")
        collectionView.delegate = collectionViewDelegateAndDataSource
        collectionView.dataSource = collectionViewDelegateAndDataSource
    }
    
    
    func setupHeaderAndTitleLabel() {
        // Initialize views and add them to the ViewController's view
        headerView = UIView()
        headerView.backgroundColor = .darkGray
        self.view.addSubview(headerView)
        
        titleLabel = UILabel()
        titleLabel.text = "LEDS"
        titleLabel.textAlignment = .center
        titleLabel.font = UIFont(name: titleLabel.font.fontName, size: 20)
        headerView.addSubview(titleLabel)
        
        // Set position of views using constraints
        headerView.translatesAutoresizingMaskIntoConstraints = false
        headerView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor).isActive = true
        headerView.topAnchor.constraint(equalTo: self.view.topAnchor).isActive = true
        headerView.widthAnchor.constraint(equalTo: self.view.widthAnchor, multiplier: 1).isActive = true
        headerView.heightAnchor.constraint(equalTo: self.view.heightAnchor, multiplier: 0.1).isActive = true
        
        titleLabel.translatesAutoresizingMaskIntoConstraints = false
        titleLabel.centerXAnchor.constraint(equalTo: headerView.centerXAnchor).isActive = true
        titleLabel.bottomAnchor.constraint(equalTo: headerView.bottomAnchor).isActive = true
        titleLabel.widthAnchor.constraint(equalTo: headerView.widthAnchor, multiplier: 0.4).isActive = true
        titleLabel.heightAnchor.constraint(equalTo: headerView.heightAnchor, multiplier: 0.5).isActive = true
    }
    
}
