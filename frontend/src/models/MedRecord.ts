import { TreatmentRecordInterface } from "./TreatmentRecord";
import { UserInterface } from "./User";
import { MedicalProductInterface } from "./MedicalProduct";

export interface MedRecordInterface{
    Amount: number,

    TreatmentRecordID: number,
    TreatmentRecord: TreatmentRecordInterface,

    UserID : number,
    User: UserInterface

    MedicalProductID: number,
    MedicalProduct: MedicalProductInterface,
}