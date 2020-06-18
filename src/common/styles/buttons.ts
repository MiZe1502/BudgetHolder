import { css } from 'emotion'

import colors from "./colors";

export const ButtonReset = css`
    border: none;
    background: transparent;
    text-decoration: none;
    outline: none;
    padding: 0;
`;

export const ButtonDefault = css`
	background-color: ${colors.greenPrimary}; 
	fill: ${colors.greenPrimary}; 
	transition: background-color 300ms ease;
    cursor: pointer;

	&:hover,
	&:active {
		background-color: ${colors.greenActive}; 
		fill: ${colors.greenActive}; 
        cursor: pointer;
	}
    
	&:active {
		background-color: ${colors.greenLight};
		fill: ${colors.greenLight};
	}

    &:disabled {
		background-color: ${colors.darkGray};
		fill: ${colors.darkGray};
        border: none;
        cursor: default;
        box-shadow: none;
        color: whitesmoke;
	}
`;