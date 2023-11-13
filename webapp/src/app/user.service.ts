import { Injectable } from '@angular/core';
import { User } from './user.model';
import {environment} from '.././environment';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root',
})
export class UserService {

  constructor(private http: HttpClient) { }

  private users: User[] = [];

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(`${environment.serverUrl}/all`)
  }

  createUser(user: User): Observable<User> {
    return this.http.post<User>(`${environment.serverUrl}/all`, user)
  }

  updateUser(user: User): void {
    const index = this.users.findIndex((u) => u.id === user.id);
    if (index !== -1) {
      this.users[index] = user;
    }
  }

  deleteUser(userId: number): void {
    this.users = this.users.filter((user) => user.id !== userId);
  }
}
