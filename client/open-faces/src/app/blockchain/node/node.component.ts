import { Component, EventEmitter, Input, OnInit, Output, SimpleChanges } from '@angular/core';
import { Node } from './node';

@Component({
  selector: 'app-node',
  templateUrl: './node.component.html',
  styleUrls: ['./node.component.css']
})
export class NodeComponent implements OnInit {
	@Input() node:Node;
	@Output() updateNodeBool = new EventEmitter<Node>();
	@Output() mineNodeBool = new EventEmitter<Node>();

	isMobile: boolean

  constructor() { }
  ngOnInit() { 
	 this.isMobile = window.innerWidth <= 400
	}

	mine() {
		this.mineNodeBool.emit(this.node);
	}

	update(event: any) {
		// ngModel only returns strings. Therefore we have to clarify data types
		// before passing over to our server for processing. 
		if (this.node.IsMined) {
			this.node.Idx = Number(this.node.Idx)
			this.node.Nonce = Number(this.node.Nonce)

			this.updateNodeBool.emit(this.node);
		}
	}
}
