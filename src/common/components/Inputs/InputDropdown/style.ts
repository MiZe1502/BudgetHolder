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
  width: 100%;
`;

export const Dropdown = css`
  position: absolute;
  top: 54px;
  max-height: ${maxHeight}px;
  width: 100%;
  background-color: ${colors.white};
  box-shadow: rgba(0, 0, 0, 0.16) 0px 4px 20px 0px;
`;

export const SingleLine = css`
  cursor: pointer;
  padding: 4px 8px;
  height: 30px;
  border-bottom: 1px solid ${colors.greenPrimary};
  
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

export const ButtonWrapper = css`
  position: absolute;
  right: 8px;
`;

export const IconButton = css`
  padding-top: 0px;
`;

export const ArrowIcon = css`
  & > svg {
    transform: rotate(90deg);
  }
`;

export const ArrowIconExpanded = css`
  & > svg {
    transform: rotate(270deg);
  }
`;