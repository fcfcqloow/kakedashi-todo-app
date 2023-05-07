export interface TaskJson {
    id: string;
    value: string;
    deadLine: number;
    topics: string[];
    priority?: number;
    done?: number;
    created: number;
}