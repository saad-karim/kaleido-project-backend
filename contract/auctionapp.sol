pragma solidity ^0.4.2;

contract auction {

    uint public startingBid;

    uint public bidCount;

    uint public highestBid;

    address public highestBidder;

    Asset public asset;

    struct Bid {
        uint auctionID;
        address bidder;
        uint amount;
    }

    struct Asset {
        string name;
        address owner;
    }

    mapping(uint => Bid) public bids;

    mapping(address => bool) public bidders;

    function auction(string assetForSale, uint bid) public {
        asset = Asset(assetForSale, msg.sender);
        startingBid = bid;
    }

    function placeBid(uint auctionID, uint amount) public {
        bids[bidCount++] = Bid(auctionID, msg.sender, amount);

        if (amount > highestBid) {
            highestBid = amount;
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