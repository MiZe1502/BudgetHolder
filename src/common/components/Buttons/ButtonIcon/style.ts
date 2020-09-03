import {css} from 'emotion';

import {defaultButtonIconSize} from "../../../styles/buttons";

export {
    ButtonReset,
    ButtonDefault,
    ButtonPrimary,
    ButtonSecondary,
    ButtonNegative,
    defaultButtonIconSize
} from "../../../styles/buttons";

export const Button = css`
    background-color: transparent;

    &:hover,
	&:active,
    &:disabled {
		background-color: transparent;
	}
`;