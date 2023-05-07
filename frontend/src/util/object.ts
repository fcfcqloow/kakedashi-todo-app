export const copyObjectKV = (to: Record<string, any>, from: any) =>
    Object.entries(from).forEach(([key, value]) => to[key] = value);