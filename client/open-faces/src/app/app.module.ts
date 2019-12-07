import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule }    from '@angular/common/http';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BlockchainComponent } from './blockchain/blockchain.component';
import { NodeComponent } from './blockchain/node/node.component';

@NgModule({
  declarations: [
    AppComponent,
    BlockchainComponent,
    NodeComponent
  ],
  imports: [
    BrowserModule,
		FormsModule,
		HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
