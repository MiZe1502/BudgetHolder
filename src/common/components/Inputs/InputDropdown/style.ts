import {css} from 'emotion'

export {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../../styles/fonts";
export {FlexVert, FlexHorCenter} from "../../../styles/positioning";
export {InvalidInput, Input} from "../../../styles/inputs";
import colors from "../../../styles/colors";

export const maxHeight = 400;

export const Wrapper = css`
  width: 100%;
  position: relative;
`;

export const DropdownInput = css`
  cursor: pointer;
`;

export const Dropdown = css`
  position: absolute;
  top: 60px;
  max-height: ${maxHeight}px;
  width: 100%;
  background-color: ${colors.white};
  box-shadow: rgba(0, 0, 0, 0.16) 0px 4px 20px 0px;
`;

export const SingleLine = css`
  cursor: pointer;
  padding: 4px 8px;
  height: 30px;
  border-bottom: 1px solid ${colors.darkGray};
  
  &:hover {
    background-color: ${colors.greenLight};
  }
  
  &:last-of-type {
    border: 0;
  }
`;

export const SelectedLine = css`
  background-color: ${colors.greenActive};
`;