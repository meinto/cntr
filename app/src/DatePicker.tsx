import React from 'react'
import { DateRangePicker, isInclusivelyBeforeDay } from 'react-dates'
import moment from 'moment'

type DatePickerProps = {
  startDate: moment.Moment;
  endDate: moment.Moment;
  setStartDate: React.Dispatch<React.SetStateAction<moment.Moment>>;
  setEndDate: React.Dispatch<React.SetStateAction<moment.Moment>>;
  focusedInput: 'startDate' | 'endDate' | null;
  setFocusedInput: React.Dispatch<React.SetStateAction<DatePickerProps['focusedInput']>>;
}
export const DatePicker = (props: DatePickerProps) => (
  <DateRangePicker
    noBorder
    hideKeyboardShortcutsPanel
    startDate={props.startDate} 
    startDateId="start-date" 
    endDate={props.endDate} 
    endDateId="end-date" 
    onDatesChange={({ startDate, endDate }) => {
      if (startDate) props.setStartDate(startDate)
      if (endDate) props.setEndDate(endDate)
    }} 
    focusedInput={props.focusedInput} 
    onFocusChange={focusedInput => props.setFocusedInput(focusedInput)} 
    isOutsideRange={day => !isInclusivelyBeforeDay(day, moment())}
    displayFormat={() => 'DD.MM.YYYY'}
  />
)