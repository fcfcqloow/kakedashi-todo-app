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

  async listDate(args?: { year?: string, month?: string }): Promise<string[]> {
    const response = (await axios.get(`${this.baseURL}/logs`, {
      params : {
        year  : args?.year || undefined,
        month : args?.month || undefined,
      },
    }));

    return Promise.resolve(response.data);
  }
}