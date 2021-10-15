import { PatientInterface } from "./Patient";

export interface ScreeningRecordInterface {
    ID: number, 
    Illnesses: string,
    Datail: string,
    Queue: string,
    PatientID: number,
    Patient: PatientInterface,
}