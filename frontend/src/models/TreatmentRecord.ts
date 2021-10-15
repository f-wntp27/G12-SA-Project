import { ScreeningRecordInterface } from "./Screening";

export interface TreatmentRecordInterface{
    ID: number,
    ToothNumber: number,
    Date: Date,
    PresciptionRaw  : string,
	Presciptionnote : string,
    ScreeningRecordID: number,
    ScreeningRecord: ScreeningRecordInterface,
}