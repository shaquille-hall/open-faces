import { Component, OnInit, Input } from '@angular/core';
import { BlockchainService } from './blockchain.service';
import { Node } from './node/node';

@Component({
  selector: 'app-blockchain',
  templateUrl: './blockchain.component.html',
  styleUrls: ['./blockchain.component.css']
})
export class BlockchainComponent implements OnInit {
  displayEmptyNodeBool: boolean	
	nodes: Node[];
	peerNodes: Node[];

  constructor(private blockchainService: BlockchainService) { }

  ngOnInit(): void {
		this.nodes = [];
		this.createNewBlockchain();
  }

	createNewBlockchain() {
		this.displayEmptyNode()
		this.blockchainService.createBlockchain().subscribe(data => {
			this.nodes = <Node[]>data
		})
	} 

	displayEmptyNode() {
		this.displayEmptyNodeBool = true
	}
		
	addNode() {
	// Allow users to add a new node to the blockchain. Mining is handled separately

		this.displayEmptyNode()
		console.log("add new")
		var node = <Node>{
			Idx: this.nodes.length,
			Data: "",
  	 	Nonce: 0,
   		HashOfPrev: "",
   		Hash: "",
  		IsValid: true,
			IsMined: false
		}
		this.nodes.push(node)
	}

	onMine(node: Node) {
		this.blockchainService.mineNode(node).subscribe(data => {
			this.nodes = <Node[]>data
		})

		this.blockchainService.getPeerNodes().subscribe(data => {
			this.peerNodes = <Node[]>data
		})
	}

	onUpdate(node: Node) {
		this.blockchainService.updateNode(node).subscribe(data => {
			this.nodes = <Node[]>data
		})
	}
}

