import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-todo-list',
  imports: [],
  templateUrl: './todo-list.html',
  styleUrl: './todo-list.css',
})
export class TodoList {
  // Recieve information/data from the parent
  @Input() todoList: string[] = [];
  
  // Setting up an event to send an output information to the app or parent
  @Output() todoDeleted = new EventEmitter<number>();


  delete(index: number){
    this.todoDeleted.emit(index);
  }

}
