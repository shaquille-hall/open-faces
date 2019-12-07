import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { finalize, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BlockchainService {
	readonly SimpleApiRecognizer = 'action' 
	private blockchainEndpoint = 'https://us-central1-open-faces.cloudfunctions.net/BuildBlockchain/';

  constructor(private http: HttpClient) { }
	
	addNewNode(newNode)  {
		 return this.http.post(this.blockchainEndpoint + "add", '{"info":' + JSON.stringify(newNode) + '}')
	}

	createBlockchain() {
		return this.http.get(this.blockchainEndpoint + "create")
	}
	
	mineNode(minedNode) {
		return this.http.post(this.blockchainEndpoint + "mine", '{"info":' + JSON.stringify(minedNode) + ' }')
	}

	getPeerNodes() {
		return this.http.get(this.blockchainEndpoint + "peerNodes")
	}

	updateNode(updatedNode) {
		return this.http.post(this.blockchainEndpoint + "update", '{"info":' + JSON.stringify(updatedNode) + '}')
	}
}


