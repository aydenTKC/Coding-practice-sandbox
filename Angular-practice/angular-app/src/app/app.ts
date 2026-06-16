import { Component, signal} from '@angular/core';
import { AddItem } from './add-item/add-item';
import { TodoList } from './todo-list/todo-list';

@Component({
  selector: 'app-root',
  imports: [AddItem, TodoList],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected readonly title = signal('angular-app');

  // Making a varible an empty array string
  todos: string[] = []



  addTodo(newTodo: string) {
    if (newTodo) {
      this.todos = [...this.todos, newTodo]
      console.log(this.todos)
    }
  }

  handleDeletedTodo(index: number) {
    this.todos = this.todos.filter((_, i) => i !== index)
  }
  
}

