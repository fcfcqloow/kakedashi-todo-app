export interface TaskPatchRequest {
    id: string;
    value: string;
    deadLine: number;
    topic?: string;
    priority?: number;
    done?: number;
    created: number;
}