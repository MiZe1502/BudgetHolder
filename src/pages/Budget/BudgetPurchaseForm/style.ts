import {css} from 'emotion'

export {
    Font312DarkGray, Font312Black, Font312Red
} from "../../../common/styles/fonts";
export {
    SideMinorPadding, FlexVert, FlexHor, FlexHorCenter
} from "../../../common/styles/positioning";

export const maxHeight = 400;

import colors from "../../../common/styles/colors";

export const TextArea = css`
    min-height: 50px;
    height: 100%;
`;

export const ValidationBlock = css`
    width: 100%;;
`

export const Form = css`
    padding-top: 8px;
    width: 100%;
`;

export const MainFieldsWrapper = css`
    border-bottom: 1px solid ${colors.darkGray};
    box-shadow: 0 2px 4px -2px rgba(0,0,0,0.25);
    padding-bottom: 16px;
    padding-left: 32px;
    padding-right: 32px;
`

export const MinorFieldsWrapper = css`
    margin-right: 32px;
    margin-left: 32px;
    padding-top: 16px;
    padding-bottom: 16px;
    border-bottom: 1px solid ${colors.darkGray};
    
    &:last-of-type {
        border-bottom: none;
    }
`;

export const MainColumn = css`
    width: 50%;
`;

export const MinorColumn = css`
    width: 33.333%;
`;

export const NotLastColumn = css`
    margin-right: 32px;
`;

export const ButtonsWrapper = css`
    padding: 16px 32px;
    justify-content: space-between;
    box-shadow: 0 -2px 4px -2px rgba(0,0,0,0.25);
`;

export const ButtonsBlock = css`
    min-width: 30%;
    justify-content: flex-end;
`;

export const ButtonForm = css`
    margin-left: 8px;
`;

export const CounterBlock = css`
    border-radius: 4px;
    border: 1px solid ${colors.greenPrimary};
    padding: 4px;
    width: 24px;
    height: 24px;
    box-sizing: border-box;
    justify-content: center;
`;

export const GoodsItemControlBlock = css`
  justify-content: space-between;
`;