import axios from "axios";
import { ISettingsRepository } from "@/core/IF/repository";
import { Settings } from "@/core/Settings";
import { SettingsJson } from "./model/SettingsJson";
import { toBeSettings, toFeSettings } from "./adapter";

export class SettingsRepository implements ISettingsRepository {
  private baseURL: string;
  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async get(): Promise<Settings> {
    const response = await axios.get(`${this.baseURL}/settings`);
    const data = response.data as SettingsJson; 
    return toFeSettings(data);
  }

  async update(settings: Settings): Promise<void> {
    await axios.patch(`${this.baseURL}/settings`, toBeSettings(settings));
  }
};
