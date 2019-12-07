import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = "Open Faces Digital Identity";
	customBlockchainButtonText = "Create New Blockchain";
	displayBlockchain = false;

	createNewBlockchain(): void {
		this.displayBlockchain = true;
	}
}
