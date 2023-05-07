import { ILogsRepository } from "@/core/IF/repository";
import axios from "axios";

export class LogsRepository implements ILogsRepository {
  private baseURL: string;
  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async get(date: string): Promise<string[]> {
   const response = (await axios.get(`${this.baseURL}/logs/${date}`));
   return Promise.resolve(response.data);
  }

  async listDate(): Promise<string[]> {
    const response = (await axios.get(`${this.baseURL}/logs`));
    return Promise.resolve(response.data);
  }
}