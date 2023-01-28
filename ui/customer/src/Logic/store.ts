import { observable, action, makeAutoObservable } from "mobx";

export default class AppStore {
  @observable navigationState: number = 0;
  constructor() {
    makeAutoObservable(this);
  }

  @action
  setNavigationState = (state: number) => {
    this.navigationState = state;
  };
}
