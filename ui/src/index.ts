import * as FriendCard from './../pages/FriendCard';

function greeter(person) {
  return "Hello, " + person;
}

let user = "Jane User";

document.body.textContent = greeter(user);
