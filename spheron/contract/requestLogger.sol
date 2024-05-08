// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RequestLogger {
    // Event declaration
    event RequestStored(uint256 requestId, string request, uint256 timestamp);

    // Structure to hold each request along with its timestamp
    struct Request {
        string data;
        uint256 timestamp;
    }

    // State variable to store requests
    mapping(uint256 => Request) public requests;

    // Counter for generating unique request IDs
    uint256 public requestCount;

    // Function to store a request
    function storeRequest(string calldata _request) external {
        // Increment the request count to ensure each ID is unique
        requestCount += 1;

        // Create a new request struct and store it
        requests[requestCount] = Request({
            data: _request,
            timestamp: block.timestamp // Current block timestamp
        });

        // Emit the event with the request ID, request data, and timestamp
        emit RequestStored(requestCount, _request, block.timestamp);
    }
}

