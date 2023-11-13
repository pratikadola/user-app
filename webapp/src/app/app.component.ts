import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { UserService } from './user.service';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {FormsModule, FormBuilder, ReactiveFormsModule, FormGroup, FormControl} from '@angular/forms';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet,MatFormFieldModule, MatInputModule, FormControl, FormsModule, MatButtonModule, MatIconModule, FormGroup, ReactiveFormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{

  controlGroup: FormGroup;

  constructor(public service: UserService, private _formBuilder: FormBuilder) {
    this.controlGroup = new FormGroup({
      name: new FormControl(this.value),
    });
  }

  options = this._formBuilder.group({
  });

  title = 'webapp';
  value = 'Clear me';
  ngOnInit(): void {
    
    this.service.getUsers().subscribe((data) => console.log(data))
    
  }

  itemclick(  ) {
    console.log(this.value)
  }
  
}   
