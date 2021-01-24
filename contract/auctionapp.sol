pragma solidity ^0.4.2;

contract auction {

    uint public bidCount;

    uint public bid;

    address public highestBidder;

    bool public closed;

    Asset public asset;

    struct Bid {
        address bidder;
        uint amount;
    }

    struct Asset {
        string name;
        address owner;
    }

    mapping(uint => Bid) public bids;

    function auction(string assetForSale, uint startingBid) public {
        asset = Asset(assetForSale, msg.sender);
        bid = startingBid;
    }

    function placeBid(uint amount) public {
        require(closed == true, "Can't place bid, auction is closed");
        require(amount <= bid, "Bid must be higher than the current bid");
        
        bids[bidCount++] = Bid(msg.sender, amount);

        if(amount > bid) {
            bid = amount;
            highestBidder = msg.sender;
        }

        emit bidPlaced();
    }

    function closeAuction() public {
        asset.owner = highestBidder;
        closed = true

        emit auctionClosed();
    }

    event bidPlaced();

    event auctionClosed();
}