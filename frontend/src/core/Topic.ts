import type { Color } from "./Color";

export class Topic {
  public id?: string;
  public value: string;
  public created: Date;
  public done?: Date;
  public color: Color;
  constructor(args: {
    id?: string;
    value: string;
    created: Date;
    color?: Color;
    done?: Date;
  }) {
    this.value = args.value;
    this.created = args.created;
    this.color = args.color ?? '#FA5';
    this.done = args.done;
    this.id = args.id;
  }

  static topicsToNameMap(topics?: Topic[]): Record<string, string> {
    if (!topics) {
      return {};
    }
  
    return topics?.
      filter((topic) => topic.id).
      reduce((prev, curr) => ({ ...prev, [curr.id!] : curr.value }), {})
  }

  static topicsToColorMap(topics?: Topic[]): Record<string, string> {
    if (!topics) {
      return {};
    }

    return topics?.
      filter((topic) => topic.id).
      reduce((prev, curr) => ({ ...prev, [curr.id!] : curr.color }), {})
  }

  static topicsToMap(topics?: Topic[]): Record<string, Topic> {
    if (!topics) {
      return {};
    }

    return topics?.
      filter((topic) => topic.id).
      reduce((prev, curr) => ({ ...prev, [curr.id!] : curr }), {})
  }

  static topicsToIds(topics?: Topic[]): string[] {
    if (!topics) {
      return [];
    }

    return topics.
      filter((topic) => topic.id).
      map((topic) => topic.id!);
  }
};
