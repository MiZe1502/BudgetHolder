import {css} from 'emotion'

export {Font312DarkGray, Font312Black} from "../../../common/styles/fonts";
export {
    SideMinorPadding, FlexVert, FlexHor
} from "../../../common/styles/positioning";

export const maxHeight = 400;

import colors from "../../../common/styles/colors";

export const TextArea = css`
    min-height: 50px;
    height: 100%;
`;

export const Form = css`
    padding-top: 8px;
`;

export const MainFieldsWrapper = css`
    border-bottom: 1px solid ${colors.darkGray};
    padding-bottom: 32px;
`

export const MinorFieldsWrapper = css`
    padding-bottom: 8px;
    padding-top: 8px;
    border-bottom: 1px solid ${colors.darkGray};
`;

export const ButtonsWrapper = css`
    padding-bottom: 8px;
    padding-top: 8px;
`