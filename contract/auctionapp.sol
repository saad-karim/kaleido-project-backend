pragma solidity ^0.4.2;

contract auction {

    uint public bidCount;

    uint public bid;

    address public highestBidder;

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
        bids[bidCount++] = Bid(msg.sender, amount);

        if(amount > bid) {
            bid = amount;
            highestBidder = msg.sender;
        }

        emit bidPlaced();
    }

    function closeAuction() public {
        asset.owner = highestBidder;

        emit auctionClosed();
    }

    event bidPlaced();

    event auctionClosed();
}